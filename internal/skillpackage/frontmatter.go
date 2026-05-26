package skillpackage

import (
	"fmt"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// ExtractSkillMDFrontMatterYAML returns the YAML source inside the first --- ... --- block and the markdown body.
func ExtractSkillMDFrontMatterYAML(raw []byte) (fmYAML string, body string, err error) {
	text := strings.TrimPrefix(string(raw), "\ufeff")
	if strings.TrimSpace(text) == "" {
		return "", "", fmt.Errorf("SKILL.md is empty")
	}
	lines := strings.Split(text, "\n")
	if len(lines) < 2 || strings.TrimSpace(lines[0]) != "---" {
		return "", "", fmt.Errorf("SKILL.md must start with YAML front matter (---) per Agent Skills standard")
	}
	var fmLines []string
	i := 1
	for i < len(lines) {
		if strings.TrimSpace(lines[i]) == "---" {
			break
		}
		fmLines = append(fmLines, lines[i])
		i++
	}
	if i >= len(lines) {
		return "", "", fmt.Errorf("SKILL.md: front matter must end with a line containing only ---")
	}
	body = strings.Join(lines[i+1:], "\n")
	body = strings.TrimSpace(body)
	fmYAML = strings.Join(fmLines, "\n")
	return fmYAML, body, nil
}

// ParseSkillMD parses SKILL.md YAML head + body.
func ParseSkillMD(raw []byte) (*SkillManifest, string, error) {
	fmYAML, body, err := ExtractSkillMDFrontMatterYAML(raw)
	if err != nil {
		return nil, "", err
	}
	var m SkillManifest
	if err := yaml.Unmarshal([]byte(fmYAML), &m); err != nil {
		return nil, "", fmt.Errorf("SKILL.md front matter: %w", err)
	}
	normalizeManifestCompat(&m)
	return &m, body, nil
}

type skillFrontMatterExport struct {
	Name          string         `yaml:"name"`
	Description   string         `yaml:"description"`
	License       string         `yaml:"license,omitempty"`
	Compatibility string         `yaml:"compatibility,omitempty"`
	Metadata      map[string]any `yaml:"metadata,omitempty"`
	AllowedTools  string         `yaml:"allowed-tools,omitempty"`
}

// BuildSkillMD serializes SKILL.md per agentskills.io.
func BuildSkillMD(m *SkillManifest, body string) ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil manifest")
	}
	normalizeManifestCompat(m)
	fm := skillFrontMatterExport{
		Name:          strings.TrimSpace(m.Name),
		Description:   strings.TrimSpace(m.Description),
		License:       strings.TrimSpace(m.License),
		Compatibility: strings.TrimSpace(m.Compatibility),
		AllowedTools:  strings.TrimSpace(m.AllowedTools),
	}
	if len(m.Metadata) > 0 {
		fm.Metadata = m.Metadata
	}
	head, err := yaml.Marshal(&fm)
	if err != nil {
		return nil, err
	}
	s := strings.TrimSpace(string(head))
	out := "---\n" + s + "\n---\n\n" + strings.TrimSpace(body) + "\n"
	return []byte(out), nil
}

func manifestTags(m *SkillManifest) []string {
	if m == nil {
		return nil
	}
	return manifestStringList(m, "tags")
}

func versionFromMetadata(m *SkillManifest) string {
	return strings.TrimSpace(manifestStringValue(m, "version"))
}

func manifestTriggers(m *SkillManifest) []string {
	return manifestStringList(m, "triggers")
}

func manifestAliases(m *SkillManifest) []string {
	return manifestStringList(m, "aliases")
}

func manifestDomains(m *SkillManifest) []string {
	return manifestStringList(m, "domains")
}

func manifestStages(m *SkillManifest) []string {
	return manifestStringList(m, "stages")
}

func manifestTargetTypes(m *SkillManifest) []string {
	return manifestStringList(m, "target_types")
}

func manifestBundleOf(m *SkillManifest) []string {
	return manifestStringList(m, "bundle_of")
}

func manifestDependsOn(m *SkillManifest) []string {
	return manifestStringList(m, "depends_on")
}

func manifestRecommendedTools(m *SkillManifest) []string {
	return manifestStringList(m, "recommended_tools")
}

func manifestRequiredTools(m *SkillManifest) []string {
	return manifestStringList(m, "required_tools")
}

func manifestRoleHints(m *SkillManifest) []string {
	return manifestStringList(m, "role_hints")
}

func manifestAutoLoadPriority(m *SkillManifest) int {
	if m == nil {
		return 0
	}
	if v := manifestStringValue(m, "autoload_priority"); v != "" {
		var n int
		if _, err := fmt.Sscanf(v, "%d", &n); err == nil {
			return n
		}
	}
	if m.Metadata != nil {
		if raw, ok := m.Metadata["autoload_priority"]; ok {
			switch v := raw.(type) {
			case int:
				return v
			case int64:
				return int(v)
			case float64:
				return int(v)
			}
		}
	}
	return 0
}

func manifestStringValue(m *SkillManifest, key string) string {
	if m == nil {
		return ""
	}
	key = strings.TrimSpace(key)
	if key == "" {
		return ""
	}
	if m.Metadata != nil {
		if raw, ok := m.Metadata[key]; ok {
			switch v := raw.(type) {
			case string:
				return strings.TrimSpace(v)
			}
		}
	}
	switch key {
	case "version":
		return strings.TrimSpace(m.Version)
	}
	return ""
}

func manifestStringList(m *SkillManifest, key string) []string {
	if m == nil {
		return nil
	}
	key = strings.TrimSpace(key)
	if key == "" {
		return nil
	}
	if m.Metadata != nil {
		if raw, ok := m.Metadata[key]; ok {
			return dedupeStrings(anyToStrings(raw))
		}
	}
	switch key {
	case "tags":
		return dedupeStrings(m.Tags)
	case "triggers":
		return dedupeStrings(m.Triggers)
	case "aliases":
		return dedupeStrings(m.Aliases)
	case "domains":
		return dedupeStrings(m.Domains)
	case "stages":
		return dedupeStrings(m.Stages)
	case "target_types":
		return dedupeStrings(m.TargetTypes)
	case "bundle_of":
		return dedupeStrings(m.BundleOf)
	case "depends_on":
		return dedupeStrings(m.DependsOn)
	case "recommended_tools":
		return dedupeStrings(m.RecommendedTools)
	case "required_tools":
		return dedupeStrings(m.RequiredTools)
	case "role_hints":
		return dedupeStrings(m.RoleHints)
	default:
		return nil
	}
}

func normalizeManifestCompat(m *SkillManifest) {
	if m == nil {
		return
	}
	if m.Metadata == nil {
		m.Metadata = make(map[string]any)
	}
	mergeMetadataString := func(key, val string) {
		val = strings.TrimSpace(val)
		if val == "" {
			return
		}
		if existing := manifestStringValue(m, key); strings.TrimSpace(existing) != "" {
			return
		}
		m.Metadata[key] = val
	}
	mergeMetadataList := func(key string, vals []string) {
		vals = dedupeStrings(vals)
		if len(vals) == 0 {
			return
		}
		if existing := manifestStringList(m, key); len(existing) > 0 {
			return
		}
		m.Metadata[key] = vals
	}
	mergeMetadataString("version", m.Version)
	mergeMetadataList("tags", m.Tags)
	mergeMetadataList("triggers", m.Triggers)
	mergeMetadataList("aliases", m.Aliases)
	mergeMetadataList("domains", m.Domains)
	mergeMetadataList("stages", m.Stages)
	mergeMetadataList("target_types", m.TargetTypes)
	mergeMetadataList("bundle_of", m.BundleOf)
	mergeMetadataList("depends_on", m.DependsOn)
	mergeMetadataList("recommended_tools", m.RecommendedTools)
	mergeMetadataList("required_tools", m.RequiredTools)
	mergeMetadataList("role_hints", m.RoleHints)
	if _, ok := m.Metadata["autoload_priority"]; !ok && m.AutoLoadPriority != 0 {
		m.Metadata["autoload_priority"] = m.AutoLoadPriority
	}

	m.Version = versionFromMetadata(m)
	m.Tags = manifestTags(m)
	m.Triggers = manifestTriggers(m)
	m.Aliases = manifestAliases(m)
	m.Domains = manifestDomains(m)
	m.Stages = manifestStages(m)
	m.TargetTypes = manifestTargetTypes(m)
	m.BundleOf = manifestBundleOf(m)
	m.DependsOn = manifestDependsOn(m)
	m.RecommendedTools = manifestRecommendedTools(m)
	m.RequiredTools = manifestRequiredTools(m)
	m.RoleHints = manifestRoleHints(m)
	m.AutoLoadPriority = manifestAutoLoadPriority(m)

	normalizeMetadataMap(m.Metadata)
}

func normalizeMetadataMap(md map[string]any) {
	if len(md) == 0 {
		return
	}
	for key, raw := range md {
		switch v := raw.(type) {
		case []string:
			md[key] = dedupeStrings(v)
		case []any:
			md[key] = dedupeStrings(anyToStrings(v))
		case string:
			md[key] = strings.TrimSpace(v)
		}
	}
}

func anyToStrings(raw any) []string {
	switch v := raw.(type) {
	case []string:
		return dedupeStrings(v)
	case []any:
		out := make([]string, 0, len(v))
		for _, item := range v {
			if s, ok := item.(string); ok {
				s = strings.TrimSpace(s)
				if s != "" {
					out = append(out, s)
				}
			}
		}
		return dedupeStrings(out)
	default:
		return nil
	}
}

func dedupeStrings(items []string) []string {
	if len(items) == 0 {
		return nil
	}
	seen := make(map[string]string, len(items))
	out := make([]string, 0, len(items))
	for _, item := range items {
		s := strings.TrimSpace(item)
		if s == "" {
			continue
		}
		key := strings.ToLower(s)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = s
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}
