package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type report struct {
	RequestedURL           string            `json:"requested_url"`
	FinalURL               string            `json:"final_url,omitempty"`
	RedirectChain          []string          `json:"redirect_chain,omitempty"`
	Method                 string            `json:"method"`
	Status                 int               `json:"status"`
	StatusText             string            `json:"status_text"`
	ContentType            string            `json:"content_type,omitempty"`
	ContentLength          int64             `json:"content_length,omitempty"`
	Title                  string            `json:"title,omitempty"`
	Server                 string            `json:"server,omitempty"`
	XPoweredBy             string            `json:"x_powered_by,omitempty"`
	Location               string            `json:"location,omitempty"`
	SetCookieNames         []string          `json:"set_cookie_names,omitempty"`
	FrameworkHints         []string          `json:"framework_hints,omitempty"`
	SecurityHeaders        map[string]string `json:"security_headers,omitempty"`
	MissingSecurityHeaders []string          `json:"missing_security_headers,omitempty"`
	Forms                  []formSummary     `json:"forms,omitempty"`
	LinksSample            []string          `json:"links_sample,omitempty"`
	BodyPreview            string            `json:"body_preview,omitempty"`
	BodyPreviewTruncated   bool              `json:"body_preview_truncated,omitempty"`
}

type formSummary struct {
	Method string   `json:"method,omitempty"`
	Action string   `json:"action,omitempty"`
	Inputs []string `json:"inputs,omitempty"`
}

var (
	titleRe = regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`)
	formRe  = regexp.MustCompile(`(?is)<form\b([^>]*)>(.*?)</form>`)
	inputRe = regexp.MustCompile(`(?is)<(?:input|textarea|select)\b([^>]*)>`)
	linkRe  = regexp.MustCompile(`(?is)(?:href|src)=["']([^"']+)["']`)
	attrFmt = `%s\s*=\s*["']([^"']+)["']`
	spaceRe = regexp.MustCompile(`\s+`)
	ctrlRe  = regexp.MustCompile(`[\x00-\x08\x0b\x0c\x0e-\x1f]`)
)

func main() {
	var (
		urlValue        string
		method          string
		headersValue    string
		body            string
		userAgent       string
		timeoutSeconds  int
		maxBodyBytes    int
		followRedirects bool
		insecureTLS     bool
	)

	flag.StringVar(&urlValue, "url", "", "target URL")
	flag.StringVar(&method, "method", "GET", "HTTP method")
	flag.StringVar(&headersValue, "headers", "", "headers separated by newlines or ||, each in Key: Value format")
	flag.StringVar(&body, "body", "", "raw request body")
	flag.StringVar(&userAgent, "user-agent", "INNP-http-framework-test/1.0", "HTTP user agent")
	flag.IntVar(&timeoutSeconds, "timeout-seconds", 15, "request timeout in seconds")
	flag.IntVar(&maxBodyBytes, "max-body-bytes", 32768, "maximum number of response body bytes to capture")
	flag.BoolVar(&followRedirects, "follow-redirects", false, "follow HTTP redirects")
	flag.BoolVar(&insecureTLS, "insecure-tls", false, "skip TLS certificate validation")
	flag.Parse()

	urlValue = strings.TrimSpace(urlValue)
	if urlValue == "" {
		writeErrorAndExit("missing required -url")
	}
	method = strings.ToUpper(strings.TrimSpace(method))
	if method == "" {
		method = http.MethodGet
	}
	if timeoutSeconds <= 0 {
		timeoutSeconds = 15
	}
	if maxBodyBytes <= 0 {
		maxBodyBytes = 32768
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()
	if insecureTLS {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //nolint:gosec
	}

	redirectChain := make([]string, 0, 4)
	client := &http.Client{
		Timeout:   time.Duration(timeoutSeconds) * time.Second,
		Transport: transport,
	}
	if followRedirects {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return fmt.Errorf("stopped after 10 redirects")
			}
			redirectChain = append(redirectChain, req.URL.String())
			return nil
		}
	} else {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	req, err := http.NewRequest(method, urlValue, strings.NewReader(body))
	if err != nil {
		writeErrorAndExit(err.Error())
	}
	req.Header.Set("User-Agent", strings.TrimSpace(userAgent))
	applyHeaders(req, headersValue)

	resp, err := client.Do(req)
	if err != nil {
		writeErrorAndExit(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, truncated, err := readBodyPreview(resp.Body, maxBodyBytes)
	if err != nil {
		writeErrorAndExit(err.Error())
	}

	bodyPreview := normalizePreview(string(bodyBytes))
	rep := report{
		RequestedURL:           urlValue,
		FinalURL:               resp.Request.URL.String(),
		RedirectChain:          uniqueStringsKeepOrder(redirectChain),
		Method:                 method,
		Status:                 resp.StatusCode,
		StatusText:             resp.Status,
		ContentType:            strings.TrimSpace(resp.Header.Get("Content-Type")),
		ContentLength:          resp.ContentLength,
		Title:                  extractTitle(bodyPreview),
		Server:                 strings.TrimSpace(resp.Header.Get("Server")),
		XPoweredBy:             strings.TrimSpace(resp.Header.Get("X-Powered-By")),
		Location:               strings.TrimSpace(resp.Header.Get("Location")),
		SetCookieNames:         extractCookieNames(resp),
		FrameworkHints:         detectFrameworkHints(resp, bodyPreview),
		SecurityHeaders:        collectSecurityHeaders(resp.Header),
		MissingSecurityHeaders: missingSecurityHeaders(resp.Header),
		Forms:                  extractForms(bodyPreview),
		LinksSample:            extractLinks(bodyPreview, 12),
		BodyPreview:            bodyPreview,
		BodyPreviewTruncated:   truncated,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(rep); err != nil {
		writeErrorAndExit(err.Error())
	}
}

func applyHeaders(req *http.Request, raw string) {
	for _, line := range splitHeaderLines(raw) {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		name := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if name == "" || value == "" {
			continue
		}
		req.Header.Add(name, value)
	}
}

func splitHeaderLines(raw string) []string {
	raw = strings.ReplaceAll(raw, "||", "\n")
	raw = strings.ReplaceAll(raw, "\r\n", "\n")
	raw = strings.ReplaceAll(raw, "\r", "\n")
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

func readBodyPreview(r io.Reader, maxBytes int) ([]byte, bool, error) {
	limited := io.LimitReader(r, int64(maxBytes)+1)
	bodyBytes, err := io.ReadAll(limited)
	if err != nil {
		return nil, false, err
	}
	if len(bodyBytes) > maxBytes {
		return bodyBytes[:maxBytes], true, nil
	}
	return bodyBytes, false, nil
}

func normalizePreview(body string) string {
	body = ctrlRe.ReplaceAllString(body, "")
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	return body
}

func extractTitle(body string) string {
	matches := titleRe.FindStringSubmatch(body)
	if len(matches) < 2 {
		return ""
	}
	return collapseWhitespace(matches[1])
}

func extractForms(body string) []formSummary {
	matches := formRe.FindAllStringSubmatch(body, -1)
	if len(matches) == 0 {
		return nil
	}

	out := make([]formSummary, 0, len(matches))
	for _, match := range matches {
		if len(match) < 3 {
			continue
		}
		attrs := match[1]
		inner := match[2]
		form := formSummary{
			Method: strings.ToUpper(extractAttr(attrs, "method")),
			Action: extractAttr(attrs, "action"),
			Inputs: extractInputNames(inner),
		}
		if form.Method == "" {
			form.Method = http.MethodGet
		}
		out = append(out, form)
	}
	return out
}

func extractAttr(rawAttrs string, name string) string {
	re := regexp.MustCompile(fmt.Sprintf(attrFmt, regexp.QuoteMeta(name)))
	matches := re.FindStringSubmatch(rawAttrs)
	if len(matches) < 2 {
		return ""
	}
	return collapseWhitespace(matches[1])
}

func extractInputNames(body string) []string {
	matches := inputRe.FindAllStringSubmatch(body, -1)
	if len(matches) == 0 {
		return nil
	}

	names := make([]string, 0, len(matches))
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		name := extractAttr(match[1], "name")
		if name == "" {
			continue
		}
		names = append(names, name)
	}
	return uniqueStringsKeepOrder(names)
}

func extractLinks(body string, limit int) []string {
	matches := linkRe.FindAllStringSubmatch(body, -1)
	if len(matches) == 0 {
		return nil
	}

	links := make([]string, 0, len(matches))
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		link := collapseWhitespace(match[1])
		if link == "" {
			continue
		}
		links = append(links, link)
	}
	links = uniqueStringsKeepOrder(links)
	if limit > 0 && len(links) > limit {
		return links[:limit]
	}
	return links
}

func extractCookieNames(resp *http.Response) []string {
	cookies := resp.Cookies()
	if len(cookies) == 0 {
		return nil
	}

	names := make([]string, 0, len(cookies))
	for _, cookie := range cookies {
		if cookie == nil || strings.TrimSpace(cookie.Name) == "" {
			continue
		}
		names = append(names, cookie.Name)
	}
	return uniqueStringsKeepOrder(names)
}

func collectSecurityHeaders(header http.Header) map[string]string {
	interesting := []string{
		"Content-Security-Policy",
		"Strict-Transport-Security",
		"X-Frame-Options",
		"X-Content-Type-Options",
		"Referrer-Policy",
		"Permissions-Policy",
		"Cross-Origin-Opener-Policy",
		"Cross-Origin-Resource-Policy",
		"Cross-Origin-Embedder-Policy",
		"Access-Control-Allow-Origin",
	}

	out := make(map[string]string)
	for _, key := range interesting {
		value := strings.TrimSpace(header.Get(key))
		if value == "" {
			continue
		}
		out[key] = value
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func missingSecurityHeaders(header http.Header) []string {
	expected := []string{
		"Content-Security-Policy",
		"Strict-Transport-Security",
		"X-Frame-Options",
		"X-Content-Type-Options",
		"Referrer-Policy",
	}

	missing := make([]string, 0, len(expected))
	for _, key := range expected {
		if strings.TrimSpace(header.Get(key)) == "" {
			missing = append(missing, key)
		}
	}
	return missing
}

func detectFrameworkHints(resp *http.Response, body string) []string {
	headerValues := strings.ToLower(strings.Join(resp.Header.Values("X-Powered-By"), " "))
	serverValue := strings.ToLower(resp.Header.Get("Server"))
	bodyLower := strings.ToLower(body)
	cookieNames := strings.ToLower(strings.Join(extractCookieNames(resp), " "))

	hints := make([]string, 0, 12)
	add := func(label string, cond bool) {
		if cond {
			hints = append(hints, label)
		}
	}

	add("JSON API", strings.Contains(strings.ToLower(resp.Header.Get("Content-Type")), "application/json"))
	add("Express", strings.Contains(headerValues, "express") || strings.Contains(cookieNames, "connect.sid"))
	add("ASP.NET", strings.Contains(headerValues, "asp.net") || strings.Contains(cookieNames, "asp.net"))
	add("PHP", strings.Contains(headerValues, "php") || strings.Contains(cookieNames, "phpsessid"))
	add("Laravel", strings.Contains(cookieNames, "laravel_session"))
	add("Django", strings.Contains(cookieNames, "csrftoken") || strings.Contains(bodyLower, "csrfmiddlewaretoken"))
	add("Spring", strings.Contains(cookieNames, "jsessionid") || strings.Contains(bodyLower, "whitelabel error page"))
	add("Ruby on Rails", strings.Contains(cookieNames, "_session") || strings.Contains(bodyLower, "csrf-param"))
	add("React", strings.Contains(bodyLower, "data-reactroot") || strings.Contains(bodyLower, "id=\"root\""))
	add("Next.js", strings.Contains(bodyLower, "__next_data__") || strings.Contains(bodyLower, "/_next/"))
	add("Angular", strings.Contains(bodyLower, "ng-version") || strings.Contains(bodyLower, "<app-root"))
	add("Vue", strings.Contains(bodyLower, "data-v-") || strings.Contains(bodyLower, "__vue__"))
	add("WordPress", strings.Contains(bodyLower, "wp-content") || strings.Contains(bodyLower, "wp-json"))
	add("Cloudflare", resp.Header.Get("CF-Ray") != "")
	add("Vercel", resp.Header.Get("X-Vercel-Id") != "")
	add("nginx", strings.Contains(serverValue, "nginx"))
	add("Apache", strings.Contains(serverValue, "apache"))
	add("openresty", strings.Contains(serverValue, "openresty"))

	return uniqueStringsKeepOrder(hints)
}

func collapseWhitespace(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	return spaceRe.ReplaceAllString(value, " ")
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
