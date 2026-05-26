// Package skillpackage provides filesystem-backed Agent Skills layout (SKILL.md + package files)
// for HTTP admin APIs. Runtime discovery and progressive loading for agents use Eino ADK skill middleware.
package skillpackage

// SkillManifest is parsed from SKILL.md front matter (https://agentskills.io/specification.md).
type SkillManifest struct {
	Name          string         `yaml:"name"`
	Description   string         `yaml:"description"`
	License       string         `yaml:"license,omitempty"`
	Compatibility string         `yaml:"compatibility,omitempty"`
	Metadata      map[string]any `yaml:"metadata,omitempty"`
	AllowedTools  string         `yaml:"allowed-tools,omitempty"`
	// Legacy compatibility fields. They are accepted on read and normalized into
	// metadata on write so existing skill packs keep working while new writes stay
	// close to the Agent Skills canonical shape.
	Version          string   `yaml:"version,omitempty"`
	Tags             []string `yaml:"tags,omitempty"`
	Triggers         []string `yaml:"triggers,omitempty"`
	Aliases          []string `yaml:"aliases,omitempty"`
	Domains          []string `yaml:"domains,omitempty"`
	Stages           []string `yaml:"stages,omitempty"`
	RecommendedTools []string `yaml:"recommended_tools,omitempty"`
	RequiredTools    []string `yaml:"required_tools,omitempty"`
	RoleHints        []string `yaml:"role_hints,omitempty"`
	AutoLoadPriority int      `yaml:"autoload_priority,omitempty"`
	TargetTypes      []string `yaml:"target_types,omitempty"`
	BundleOf         []string `yaml:"bundle_of,omitempty"`
	DependsOn        []string `yaml:"depends_on,omitempty"`
}

// SkillSummary is API metadata for one skill directory.
type SkillSummary struct {
	ID               string         `json:"id"`
	DirName          string         `json:"dir_name"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	Version          string         `json:"version"`
	Path             string         `json:"path"`
	Tags             []string       `json:"tags"`
	Triggers         []string       `json:"triggers,omitempty"`
	Aliases          []string       `json:"aliases,omitempty"`
	Domains          []string       `json:"domains,omitempty"`
	Stages           []string       `json:"stages,omitempty"`
	TargetTypes      []string       `json:"target_types,omitempty"`
	BundleOf         []string       `json:"bundle_of,omitempty"`
	DependsOn        []string       `json:"depends_on,omitempty"`
	IsBundle         bool           `json:"is_bundle,omitempty"`
	RecommendedTools []string       `json:"recommended_tools,omitempty"`
	RequiredTools    []string       `json:"required_tools,omitempty"`
	RoleHints        []string       `json:"role_hints,omitempty"`
	AutoLoadPriority int            `json:"autoload_priority,omitempty"`
	ScriptCount      int            `json:"script_count"`
	FileCount        int            `json:"file_count"`
	FileSize         int64          `json:"file_size"`
	ModTime          string         `json:"mod_time"`
	Progressive      bool           `json:"progressive"`
	MatchScore       int            `json:"match_score,omitempty"`
	MatchReasons     []string       `json:"match_reasons,omitempty"`
	QualityScore     int            `json:"quality_score,omitempty"`
	QualityWarnings  []string       `json:"quality_warnings,omitempty"`
	Coverage         map[string]int `json:"coverage,omitempty"`
}

// SkillScriptInfo describes a file under scripts/.
type SkillScriptInfo struct {
	Name        string `json:"name"`
	RelPath     string `json:"rel_path"`
	Description string `json:"description,omitempty"`
	Size        int64  `json:"size"`
}

// SkillSection is derived from ## headings in SKILL.md.
type SkillSection struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Heading string `json:"heading"`
	Level   int    `json:"level"`
}

// PackageFileInfo describes one file inside a package.
type PackageFileInfo struct {
	Path  string `json:"path"`
	Size  int64  `json:"size"`
	IsDir bool   `json:"is_dir,omitempty"`
}

// SkillView is a loaded package for admin / API.
type SkillView struct {
	DirName          string            `json:"dir_name"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	Content          string            `json:"content"`
	Path             string            `json:"path"`
	Version          string            `json:"version"`
	Tags             []string          `json:"tags"`
	Triggers         []string          `json:"triggers,omitempty"`
	Aliases          []string          `json:"aliases,omitempty"`
	Domains          []string          `json:"domains,omitempty"`
	Stages           []string          `json:"stages,omitempty"`
	TargetTypes      []string          `json:"target_types,omitempty"`
	BundleOf         []string          `json:"bundle_of,omitempty"`
	DependsOn        []string          `json:"depends_on,omitempty"`
	IsBundle         bool              `json:"is_bundle,omitempty"`
	RecommendedTools []string          `json:"recommended_tools,omitempty"`
	RequiredTools    []string          `json:"required_tools,omitempty"`
	RoleHints        []string          `json:"role_hints,omitempty"`
	AutoLoadPriority int               `json:"autoload_priority,omitempty"`
	Metadata         map[string]any    `json:"metadata,omitempty"`
	AllowedTools     string            `json:"allowed_tools,omitempty"`
	Scripts          []SkillScriptInfo `json:"scripts,omitempty"`
	Sections         []SkillSection    `json:"sections,omitempty"`
	PackageFiles     []PackageFileInfo `json:"package_files,omitempty"`
	QualityScore     int               `json:"quality_score,omitempty"`
	QualityWarnings  []string          `json:"quality_warnings,omitempty"`
	Coverage         map[string]int    `json:"coverage,omitempty"`
}

// SkillRecommendOptions tunes skill ranking for one request / search intent.
type SkillRecommendOptions struct {
	Query          string
	Role           string
	AvailableTools []string
	Limit          int
}

// SkillRecommendation is a ranked recommendation result for one skill package.
type SkillRecommendation struct {
	ID               string   `json:"id"`
	DirName          string   `json:"dir_name"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Version          string   `json:"version"`
	Tags             []string `json:"tags,omitempty"`
	Triggers         []string `json:"triggers,omitempty"`
	Aliases          []string `json:"aliases,omitempty"`
	Domains          []string `json:"domains,omitempty"`
	Stages           []string `json:"stages,omitempty"`
	TargetTypes      []string `json:"target_types,omitempty"`
	BundleOf         []string `json:"bundle_of,omitempty"`
	DependsOn        []string `json:"depends_on,omitempty"`
	IsBundle         bool     `json:"is_bundle,omitempty"`
	BundleSources    []string `json:"bundle_sources,omitempty"`
	MatchedTargets   []string `json:"matched_targets,omitempty"`
	RecommendedTools []string `json:"recommended_tools,omitempty"`
	RequiredTools    []string `json:"required_tools,omitempty"`
	RoleHints        []string `json:"role_hints,omitempty"`
	AutoLoadPriority int      `json:"autoload_priority,omitempty"`
	Score            int      `json:"score"`
	Reasons          []string `json:"reasons,omitempty"`
	QualityScore     int      `json:"quality_score,omitempty"`
	QualityWarnings  []string `json:"quality_warnings,omitempty"`
}

// SkillRecommendResult captures routing context plus ranked bundle/skill suggestions.
type SkillRecommendResult struct {
	Query           string                `json:"query,omitempty"`
	Role            string                `json:"role,omitempty"`
	AvailableTools  []string              `json:"available_tools,omitempty"`
	PrimaryTarget   string                `json:"primary_target,omitempty"`
	InferredTargets []string              `json:"inferred_targets,omitempty"`
	Bundles         []SkillRecommendation `json:"bundles,omitempty"`
	Recommendations []SkillRecommendation `json:"recommendations"`
}
