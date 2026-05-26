package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type analysis struct {
	TokenType    string             `json:"token_type"`
	Segments     int                `json:"segments"`
	Algorithm    string             `json:"algorithm,omitempty"`
	KeyID        string             `json:"kid,omitempty"`
	Header       map[string]any     `json:"header,omitempty"`
	Payload      map[string]any     `json:"payload,omitempty"`
	HeaderRaw    string             `json:"header_raw,omitempty"`
	PayloadRaw   string             `json:"payload_raw,omitempty"`
	Findings     []string           `json:"findings,omitempty"`
	TimeChecks   map[string]any     `json:"time_checks,omitempty"`
	Verification *verificationCheck `json:"verification,omitempty"`
	ParseError   string             `json:"parse_error,omitempty"`
}

type verificationCheck struct {
	Algorithm          string `json:"algorithm,omitempty"`
	CheckedSecrets     int    `json:"checked_secrets,omitempty"`
	SignatureValid     bool   `json:"signature_valid"`
	MatchedSecretLabel string `json:"matched_secret_label,omitempty"`
	SkippedReason      string `json:"skipped_reason,omitempty"`
}

func main() {
	var (
		token            string
		secret           string
		secretCandidates string
	)

	flag.StringVar(&token, "token", "", "JWT or bearer token value")
	flag.StringVar(&secret, "secret", "", "single HMAC secret candidate")
	flag.StringVar(&secretCandidates, "secret-candidates", "", "multiple HMAC secret candidates separated by commas or newlines")
	flag.Parse()

	token = normalizeToken(token)
	if token == "" {
		writeErrorAndExit("missing required -token")
	}

	result := analyzeToken(token, secret, secretCandidates)
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		writeErrorAndExit(err.Error())
	}
}

func analyzeToken(token, secret, secretCandidates string) analysis {
	parts := strings.Split(token, ".")
	result := analysis{
		TokenType: "unknown",
		Segments:  len(parts),
		Findings:  make([]string, 0, 8),
	}

	switch len(parts) {
	case 3:
		result.TokenType = "jwt"
	case 5:
		result.TokenType = "jwe_like"
		result.Findings = append(result.Findings, "token has 5 segments and looks like JWE; payload decryption is not performed")
	default:
		result.ParseError = fmt.Sprintf("unexpected token format with %d segments", len(parts))
		return result
	}

	headerBytes, err := decodeSegment(parts[0])
	if err != nil {
		result.ParseError = fmt.Sprintf("failed to decode header: %v", err)
		return result
	}
	result.HeaderRaw = string(headerBytes)

	headerMap, err := decodeJSONObject(headerBytes)
	if err != nil {
		result.ParseError = fmt.Sprintf("header is not valid JSON: %v", err)
		return result
	}
	result.Header = headerMap
	result.Algorithm = strings.TrimSpace(toString(headerMap["alg"]))
	result.KeyID = strings.TrimSpace(toString(headerMap["kid"]))

	if len(parts) == 5 {
		addHeaderFindings(&result, headerMap)
		return result
	}

	payloadBytes, err := decodeSegment(parts[1])
	if err != nil {
		result.ParseError = fmt.Sprintf("failed to decode payload: %v", err)
		addHeaderFindings(&result, headerMap)
		return result
	}
	result.PayloadRaw = string(payloadBytes)

	payloadMap, err := decodeJSONObject(payloadBytes)
	if err != nil {
		result.ParseError = fmt.Sprintf("payload is not valid JSON: %v", err)
		addHeaderFindings(&result, headerMap)
		return result
	}
	result.Payload = payloadMap

	addHeaderFindings(&result, headerMap)
	addPayloadFindings(&result, payloadMap)
	result.TimeChecks = buildTimeChecks(payloadMap, &result.Findings)

	result.Verification = verifyHMACSignature(parts, result.Algorithm, secret, secretCandidates)
	if result.Verification != nil && result.Verification.SignatureValid {
		result.Findings = append(result.Findings, "signature matched one of the provided HMAC secrets")
	}

	result.Findings = uniqueStringsKeepOrder(result.Findings)
	return result
}

func addHeaderFindings(result *analysis, header map[string]any) {
	if result == nil {
		return
	}
	alg := strings.ToUpper(strings.TrimSpace(toString(header["alg"])))
	switch alg {
	case "":
		result.Findings = append(result.Findings, "missing alg header")
	case "NONE":
		result.Findings = append(result.Findings, "alg=none indicates unsigned token acceptance should be checked")
	}

	for _, key := range []string{"kid", "jku", "jwk", "x5u", "x5c", "crit"} {
		if _, ok := header[key]; ok {
			result.Findings = append(result.Findings, fmt.Sprintf("header contains %s and may influence key selection or validation logic", key))
		}
	}
}

func addPayloadFindings(result *analysis, payload map[string]any) {
	if result == nil {
		return
	}

	if _, ok := payload["exp"]; !ok {
		result.Findings = append(result.Findings, "missing exp claim")
	}
	if role := strings.ToLower(strings.TrimSpace(toString(payload["role"]))); role == "admin" {
		result.Findings = append(result.Findings, "payload claims admin role")
	}
	if admin, ok := payload["admin"].(bool); ok && admin {
		result.Findings = append(result.Findings, "payload contains admin=true")
	}
	if len(toString(payload["aud"])) == 0 {
		result.Findings = append(result.Findings, "missing aud claim")
	}
}

func buildTimeChecks(payload map[string]any, findings *[]string) map[string]any {
	now := time.Now().Unix()
	checks := map[string]any{
		"now_unix": now,
	}

	iat, hasIat := claimAsUnix(payload["iat"])
	exp, hasExp := claimAsUnix(payload["exp"])
	nbf, hasNbf := claimAsUnix(payload["nbf"])

	if hasIat {
		checks["iat_unix"] = iat
		checks["iat_rfc3339"] = time.Unix(iat, 0).UTC().Format(time.RFC3339)
	}
	if hasExp {
		checks["exp_unix"] = exp
		checks["exp_rfc3339"] = time.Unix(exp, 0).UTC().Format(time.RFC3339)
		checks["expired"] = now > exp
		if now > exp && findings != nil {
			*findings = append(*findings, "token is expired")
		}
	}
	if hasNbf {
		checks["nbf_unix"] = nbf
		checks["nbf_rfc3339"] = time.Unix(nbf, 0).UTC().Format(time.RFC3339)
		checks["not_yet_valid"] = now < nbf
		if now < nbf && findings != nil {
			*findings = append(*findings, "token is not yet valid")
		}
	}
	if hasExp && hasIat && exp > iat && exp-iat > 30*24*60*60 {
		checks["long_lived"] = true
		if findings != nil {
			*findings = append(*findings, "token lifetime is longer than 30 days")
		}
	}
	return checks
}

func verifyHMACSignature(parts []string, algorithm, secret, secretCandidates string) *verificationCheck {
	result := &verificationCheck{
		Algorithm: strings.ToUpper(strings.TrimSpace(algorithm)),
	}
	if len(parts) != 3 {
		result.SkippedReason = "verification requires a 3-segment JWS token"
		return result
	}

	candidates := buildSecretCandidates(secret, secretCandidates)
	if len(candidates) == 0 {
		result.SkippedReason = "no HMAC secret candidates were provided"
		return result
	}

	signingInput := parts[0] + "." + parts[1]
	signature, err := decodeSegment(parts[2])
	if err != nil {
		result.SkippedReason = fmt.Sprintf("signature decode failed: %v", err)
		return result
	}

	var macFactory func() []byte
	switch result.Algorithm {
	case "HS256":
		macFactory = func() []byte { return hmacDigestSHA256(signingInput, "") }
	case "HS384":
		macFactory = func() []byte { return hmacDigestSHA384(signingInput, "") }
	case "HS512":
		macFactory = func() []byte { return hmacDigestSHA512(signingInput, "") }
	default:
		result.SkippedReason = "signature verification is currently limited to HS256/HS384/HS512"
		return result
	}

	for _, candidate := range candidates {
		result.CheckedSecrets++

		var expected []byte
		switch result.Algorithm {
		case "HS256":
			expected = hmacDigestSHA256(signingInput, candidate.Value)
		case "HS384":
			expected = hmacDigestSHA384(signingInput, candidate.Value)
		case "HS512":
			expected = hmacDigestSHA512(signingInput, candidate.Value)
		default:
			expected = macFactory()
		}

		if hmac.Equal(expected, signature) {
			result.SignatureValid = true
			result.MatchedSecretLabel = candidate.Label
			return result
		}
	}

	return result
}

type labeledSecret struct {
	Label string
	Value string
}

func buildSecretCandidates(secret, raw string) []labeledSecret {
	out := make([]labeledSecret, 0, 4)
	add := func(label string, value string) {
		value = strings.TrimSpace(value)
		if value == "" {
			return
		}
		out = append(out, labeledSecret{Label: label, Value: value})
	}

	add("secret", secret)
	for idx, part := range splitSecretList(raw) {
		add(fmt.Sprintf("secret_candidates[%d]", idx), part)
	}
	return out
}

func splitSecretList(raw string) []string {
	raw = strings.ReplaceAll(raw, "\r\n", "\n")
	raw = strings.ReplaceAll(raw, "\r", "\n")
	raw = strings.ReplaceAll(raw, ",", "\n")
	lines := strings.Split(raw, "\n")
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		out = append(out, line)
	}
	return out
}

func claimAsUnix(value any) (int64, bool) {
	switch v := value.(type) {
	case float64:
		return int64(v), true
	case int64:
		return v, true
	case int:
		return int64(v), true
	case string:
		if strings.TrimSpace(v) == "" {
			return 0, false
		}
		var parsed int64
		if _, err := fmt.Sscanf(strings.TrimSpace(v), "%d", &parsed); err == nil {
			return parsed, true
		}
	}
	return 0, false
}

func normalizeToken(token string) string {
	token = strings.TrimSpace(token)
	token = strings.TrimPrefix(token, "Bearer ")
	token = strings.TrimPrefix(token, "bearer ")
	return strings.TrimSpace(token)
}

func decodeSegment(segment string) ([]byte, error) {
	segment = strings.TrimSpace(segment)
	if segment == "" {
		return nil, fmt.Errorf("segment is empty")
	}
	if mod := len(segment) % 4; mod != 0 {
		segment += strings.Repeat("=", 4-mod)
	}
	return base64.URLEncoding.DecodeString(segment)
}

func decodeJSONObject(raw []byte) (map[string]any, error) {
	var out map[string]any
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func toString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	default:
		if value == nil {
			return ""
		}
		return fmt.Sprintf("%v", value)
	}
}

func hmacDigestSHA256(data, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(data))
	return mac.Sum(nil)
}

func hmacDigestSHA384(data, secret string) []byte {
	mac := hmac.New(sha512.New384, []byte(secret))
	_, _ = mac.Write([]byte(data))
	return mac.Sum(nil)
}

func hmacDigestSHA512(data, secret string) []byte {
	mac := hmac.New(sha512.New, []byte(secret))
	_, _ = mac.Write([]byte(data))
	return mac.Sum(nil)
}

func uniqueStringsKeepOrder(items []string) []string {
	if len(items) == 0 {
		return nil
	}

	seen := make(map[string]struct{}, len(items))
	out := make([]string, 0, len(items))
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		key := strings.ToLower(item)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, item)
	}
	return out
}

func writeErrorAndExit(message string) {
	payload := map[string]string{"error": strings.TrimSpace(message)}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(payload)
	os.Exit(1)
}
