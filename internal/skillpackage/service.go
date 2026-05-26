package skillpackage

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

type skillPackageDoc struct {
	summary         SkillSummary
	manifest        *SkillManifest
	body            string
	scripts         []SkillScriptInfo
	sections        []SkillSection
	packageFiles    []PackageFileInfo
	searchFields    []string
	qualityScore    int
	qualityWarnings []string
	coverage        map[string]int
}

type targetProfile struct {
	Name      string
	Keywords  []string
	RoleHints []string
	ToolHints []string
}

var defaultTargetProfiles = []targetProfile{
	{
		Name:      "ctf",
		Keywords:  []string{"ctf", "challenge", "flag", "misc", "jeopardy"},
		RoleHints: []string{"ctf", "pwn", "reverse"},
		ToolHints: []string{"ghidra", "radare2", "gdb", "pwntools", "binwalk", "volatility3"},
	},
	{
		Name:      "web",
		Keywords:  []string{"web", "http", "https", "browser", "cookie", "jwt", "session", "xss", "ssti", "ssrf", "csrf"},
		RoleHints: []string{"web", "application"},
		ToolHints: []string{"http-framework-test", "ffuf", "sqlmap", "nuclei", "jwt-analyzer"},
	},
	{
		Name:      "api",
		Keywords:  []string{"api", "openapi", "swagger", "graphql", "rest", "endpoint", "contract"},
		RoleHints: []string{"api", "backend"},
		ToolHints: []string{"http-framework-test", "nuclei", "jwt-analyzer"},
	},
	{
		Name:      "binary",
		Keywords:  []string{"elf", "pe", "binary", "heap", "stack", "rop", "libc", "shellcode", "nc"},
		RoleHints: []string{"binary", "exploit", "pwn"},
		ToolHints: []string{"ghidra", "radare2", "gdb", "pwntools", "strings"},
	},
	{
		Name:      "reverse",
		Keywords:  []string{"reverse", "decompile", "disasm", "wasm", "apk", "firmware", "vm"},
		RoleHints: []string{"reverse", "re"},
		ToolHints: []string{"ghidra", "radare2", "strings", "objdump"},
	},
	{
		Name:      "forensics",
		Keywords:  []string{"pcap", "memory", "registry", "stego", "image", "disk", "forensics", "metadata"},
		RoleHints: []string{"forensics", "incident response"},
		ToolHints: []string{"volatility3", "binwalk", "exiftool", "foremost", "zsteg"},
	},
	{
		Name:      "crypto",
		Keywords:  []string{"rsa", "aes", "ecc", "hash", "signature", "nonce", "cipher", "crypto", "padding"},
		RoleHints: []string{"crypto", "cryptography"},
		ToolHints: []string{"hashcat", "john"},
	},
	{
		Name:      "cloud",
		Keywords:  []string{"cloud", "k8s", "kubernetes", "container", "docker", "iam", "s3", "lambda"},
		RoleHints: []string{"cloud", "devsecops"},
		ToolHints: []string{"nuclei"},
	},
}

// ListSkillSummaries scans skillsRoot and returns index rows for the admin API.
func ListSkillSummaries(skillsRoot string) ([]SkillSummary, error) {
	docs, err := loadSkillDocs(skillsRoot)
	if err != nil {
		return nil, err
	}
	out := make([]SkillSummary, 0, len(docs))
	for _, doc := range docs {
		out = append(out, doc.summary)
	}
	return out, nil
}

// LoadOptions mirrors legacy API query params for the web admin.
type LoadOptions struct {
	Depth   string // summary | full
	Section string
}

// LoadSkill returns manifest + body + package listing for admin.
func LoadSkill(skillsRoot, skillID string, opt LoadOptions) (*SkillView, error) {
	doc, err := loadSkillDoc(skillsRoot, skillID)
	if err != nil {
		return nil, err
	}
	v := &SkillView{
		DirName:          doc.summary.DirName,
		Name:             doc.summary.Name,
		Description:      doc.summary.Description,
		Content:          doc.body,
		Path:             doc.summary.Path,
		Version:          doc.summary.Version,
		Tags:             append([]string(nil), doc.summary.Tags...),
		Triggers:         append([]string(nil), doc.summary.Triggers...),
		Aliases:          append([]string(nil), doc.summary.Aliases...),
		Domains:          append([]string(nil), doc.summary.Domains...),
		Stages:           append([]string(nil), doc.summary.Stages...),
		TargetTypes:      append([]string(nil), doc.summary.TargetTypes...),
		BundleOf:         append([]string(nil), doc.summary.BundleOf...),
		DependsOn:        append([]string(nil), doc.summary.DependsOn...),
		IsBundle:         doc.summary.IsBundle,
		RecommendedTools: append([]string(nil), doc.summary.RecommendedTools...),
		RequiredTools:    append([]string(nil), doc.summary.RequiredTools...),
		RoleHints:        append([]string(nil), doc.summary.RoleHints...),
		AutoLoadPriority: doc.summary.AutoLoadPriority,
		Metadata:         cloneMetadata(doc.manifest.Metadata),
		AllowedTools:     strings.TrimSpace(doc.manifest.AllowedTools),
		Scripts:          append([]SkillScriptInfo(nil), doc.scripts...),
		Sections:         append([]SkillSection(nil), doc.sections...),
		PackageFiles:     append([]PackageFileInfo(nil), doc.packageFiles...),
		QualityScore:     doc.qualityScore,
		QualityWarnings:  append([]string(nil), doc.qualityWarnings...),
		Coverage:         cloneCoverageMap(doc.coverage),
	}
	depth := strings.ToLower(strings.TrimSpace(opt.Depth))
	if depth == "" {
		depth = "full"
	}
	sec := strings.TrimSpace(opt.Section)
	if sec != "" {
		mds := splitMarkdownSections(doc.body)
		chunk := findSectionContent(mds, sec)
		if chunk == "" {
			v.Content = fmt.Sprintf("_(section %q not found in SKILL.md for skill %s)_", sec, skillID)
		} else {
			v.Content = chunk
		}
		return v, nil
	}
	if depth == "summary" {
		v.Content = buildSummaryMarkdown(doc.summary.Name, doc.summary.Description, doc.summary.Tags, doc.scripts, doc.sections, doc.body)
	}
	return v, nil
}

// RecommendSkills ranks skills for a query / role / current tool context.
func RecommendSkills(skillsRoot string, opt SkillRecommendOptions) ([]SkillRecommendation, error) {
	result, err := RecommendSkillsDetailed(skillsRoot, opt)
	if err != nil {
		return nil, err
	}
	return result.Recommendations, nil
}

// RecommendSkillsDetailed ranks skills and returns inferred routing context.
func RecommendSkillsDetailed(skillsRoot string, opt SkillRecommendOptions) (*SkillRecommendResult, error) {
	docs, err := loadSkillDocs(skillsRoot)
	if err != nil {
		return nil, err
	}
	limit := opt.Limit
	if limit <= 0 {
		limit = 5
	}
	queryTerms := normalizeSearchTerms(opt.Query)
	roleTerms := normalizeSearchTerms(opt.Role)
	availableTools := uniqueStringsKeepOrder(opt.AvailableTools)
	availableToolSet := normalizeLowerSet(availableTools)
	inferredTargets := inferTargetTypes(opt.Query, opt.Role, availableTools)
	primaryTarget := ""
	if len(inferredTargets) > 0 {
		primaryTarget = inferredTargets[0]
	}

	hasContext := len(queryTerms) > 0 || len(roleTerms) > 0 || len(availableToolSet) > 0
	recs := make([]SkillRecommendation, 0, len(docs))
	bundles := make([]SkillRecommendation, 0, len(docs))
	for _, doc := range docs {
		rec := scoreSkillDoc(doc, queryTerms, roleTerms, availableToolSet, inferredTargets)
		if hasContext && rec.Score <= 0 {
			continue
		}
		if rec.IsBundle {
			bundles = append(bundles, rec)
			continue
		}
		recs = append(recs, rec)
	}

	sortRecommendations(bundles)
	attachBundleSignals(recs, bundles)
	sortRecommendations(recs)

	if len(bundles) > limit {
		bundles = bundles[:limit]
	}
	if len(recs) > limit {
		recs = recs[:limit]
	}

	return &SkillRecommendResult{
		Query:           strings.TrimSpace(opt.Query),
		Role:            strings.TrimSpace(opt.Role),
		AvailableTools:  availableTools,
		PrimaryTarget:   primaryTarget,
		InferredTargets: inferredTargets,
		Bundles:         bundles,
		Recommendations: recs,
	}, nil
}

func sortRecommendations(recs []SkillRecommendation) {
	sort.Slice(recs, func(i, j int) bool {
		if recs[i].Score != recs[j].Score {
			return recs[i].Score > recs[j].Score
		}
		if recs[i].AutoLoadPriority != recs[j].AutoLoadPriority {
			return recs[i].AutoLoadPriority > recs[j].AutoLoadPriority
		}
		return recs[i].DirName < recs[j].DirName
	})
}

// BuildSkillShortlistInstruction returns a compact recommendation block for model instructions.
func BuildSkillShortlistInstruction(skillsRoot string, opt SkillRecommendOptions) (string, []SkillRecommendation, error) {
	result, err := RecommendSkillsDetailed(skillsRoot, opt)
	if err != nil {
		return "", nil, err
	}
	listed := result.Recommendations
	if len(listed) == 0 {
		listed = result.Bundles
	}
	if len(listed) == 0 {
		return "", nil, nil
	}

	var sb strings.Builder
	sb.WriteString("Skill shortlist for this request. Prefer these skills before broad search.\n")
	if result.PrimaryTarget != "" || len(result.InferredTargets) > 0 {
		sb.WriteString("- route")
		if result.PrimaryTarget != "" {
			sb.WriteString(": primary_target=")
			sb.WriteString(result.PrimaryTarget)
		}
		if len(result.InferredTargets) > 0 {
			sb.WriteString(" targets=")
			sb.WriteString(strings.Join(result.InferredTargets, "/"))
		}
		sb.WriteByte('\n')
	}
	if len(result.Bundles) > 0 {
		names := make([]string, 0, len(result.Bundles))
		for i, rec := range result.Bundles {
			if i >= 2 {
				break
			}
			names = append(names, rec.Name)
		}
		if len(names) > 0 {
			sb.WriteString("- bundles: ")
			sb.WriteString(strings.Join(names, ", "))
			sb.WriteByte('\n')
		}
	}
	for _, rec := range listed {
		sb.WriteString("- ")
		sb.WriteString(rec.Name)
		if rec.Description != "" {
			sb.WriteString(": ")
			sb.WriteString(rec.Description)
		}
		var extras []string
		if len(rec.Tags) > 0 {
			extras = append(extras, "tags="+strings.Join(rec.Tags, "/"))
		}
		if len(rec.Triggers) > 0 {
			extras = append(extras, "triggers="+strings.Join(rec.Triggers, "/"))
		}
		if len(rec.MatchedTargets) > 0 {
			extras = append(extras, "targets="+strings.Join(rec.MatchedTargets, "/"))
		}
		if rec.IsBundle && len(rec.BundleOf) > 0 {
			extras = append(extras, "bundle_of="+strings.Join(rec.BundleOf, "/"))
		}
		if !rec.IsBundle && len(rec.BundleSources) > 0 {
			extras = append(extras, "bundled_by="+strings.Join(rec.BundleSources, "/"))
		}
		if len(rec.Reasons) > 0 {
			extras = append(extras, "reasons="+strings.Join(rec.Reasons, "; "))
		}
		if len(extras) > 0 {
			sb.WriteString(" [")
			sb.WriteString(strings.Join(extras, " | "))
			sb.WriteString("]")
		}
		sb.WriteByte('\n')
	}
	sb.WriteString("Rule: load a full skill only when you need its workflow, scripts, or references; do not bulk-load every skill.")
	return sb.String(), listed, nil
}

// ReadScriptText returns file content as string (for HTTP resource_path).
func ReadScriptText(skillsRoot, skillID, relPath string, maxBytes int64) (string, error) {
	b, err := ReadPackageFile(skillsRoot, skillID, relPath, maxBytes)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func loadSkillDocs(skillsRoot string) ([]skillPackageDoc, error) {
	names, err := ListSkillDirNames(skillsRoot)
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	out := make([]skillPackageDoc, 0, len(names))
	for _, dirName := range names {
		doc, err := loadSkillDoc(skillsRoot, dirName)
		if err != nil {
			continue
		}
		out = append(out, *doc)
	}
	return out, nil
}

func loadSkillDoc(skillsRoot, dirName string) (*skillPackageDoc, error) {
	skillPath := SkillDir(skillsRoot, dirName)
	mdPath, err := ResolveSKILLPath(skillPath)
	if err != nil {
		return nil, err
	}
	raw, err := os.ReadFile(mdPath)
	if err != nil {
		return nil, err
	}
	man, body, err := ParseSkillMD(raw)
	if err != nil {
		return nil, err
	}
	if err := ValidateAgentSkillManifestInPackage(man, dirName); err != nil {
		return nil, err
	}
	fi, err := os.Stat(mdPath)
	if err != nil {
		return nil, err
	}
	pfiles, err := ListPackageFiles(skillsRoot, dirName)
	if err != nil {
		return nil, err
	}
	scripts, err := listScripts(skillsRoot, dirName)
	if err != nil {
		return nil, err
	}
	sort.Slice(scripts, func(i, j int) bool { return scripts[i].RelPath < scripts[j].RelPath })
	sections := deriveSections(body)

	nFiles := 0
	for _, p := range pfiles {
		if !p.IsDir {
			nFiles++
		}
	}

	targetTypes := manifestTargetTypes(man)
	bundleOf := manifestBundleOf(man)
	dependsOn := manifestDependsOn(man)
	summary := SkillSummary{
		ID:               dirName,
		DirName:          dirName,
		Name:             man.Name,
		Description:      man.Description,
		Version:          versionFromMetadata(man),
		Path:             skillPath,
		Tags:             manifestTags(man),
		Triggers:         manifestTriggers(man),
		Aliases:          manifestAliases(man),
		Domains:          manifestDomains(man),
		Stages:           manifestStages(man),
		TargetTypes:      targetTypes,
		BundleOf:         bundleOf,
		DependsOn:        dependsOn,
		IsBundle:         len(bundleOf) > 0,
		RecommendedTools: manifestRecommendedTools(man),
		RequiredTools:    manifestRequiredTools(man),
		RoleHints:        manifestRoleHints(man),
		AutoLoadPriority: manifestAutoLoadPriority(man),
		ScriptCount:      len(scripts),
		FileCount:        nFiles,
		FileSize:         fi.Size(),
		ModTime:          fi.ModTime().Format("2006-01-02 15:04:05"),
		Progressive:      true,
	}
	qualityScore, qualityWarnings, coverage := scoreSkillQuality(man, body, scripts, sections)
	summary.QualityScore = qualityScore
	summary.QualityWarnings = append([]string(nil), qualityWarnings...)
	summary.Coverage = cloneCoverageMap(coverage)

	doc := &skillPackageDoc{
		summary:         summary,
		manifest:        man,
		body:            body,
		scripts:         scripts,
		sections:        sections,
		packageFiles:    pfiles,
		qualityScore:    qualityScore,
		qualityWarnings: append([]string(nil), qualityWarnings...),
		coverage:        cloneCoverageMap(coverage),
	}
	doc.searchFields = buildSearchFields(doc)
	return doc, nil
}

func buildSearchFields(doc *skillPackageDoc) []string {
	if doc == nil {
		return nil
	}
	fields := make([]string, 0, 48)
	add := func(s string) {
		s = strings.TrimSpace(strings.ToLower(s))
		if s != "" {
			fields = append(fields, s)
		}
	}
	add(doc.summary.DirName)
	add(doc.summary.Name)
	add(doc.summary.Description)
	add(doc.body)
	for _, item := range doc.summary.Tags {
		add(item)
	}
	for _, item := range doc.summary.Triggers {
		add(item)
	}
	for _, item := range doc.summary.Aliases {
		add(item)
	}
	for _, item := range doc.summary.Domains {
		add(item)
	}
	for _, item := range doc.summary.Stages {
		add(item)
	}
	for _, item := range doc.summary.TargetTypes {
		add(item)
	}
	for _, item := range doc.summary.BundleOf {
		add(item)
	}
	for _, item := range doc.summary.DependsOn {
		add(item)
	}
	for _, item := range doc.summary.RecommendedTools {
		add(item)
	}
	for _, item := range doc.summary.RequiredTools {
		add(item)
	}
	for _, item := range doc.summary.RoleHints {
		add(item)
	}
	for _, sec := range doc.sections {
		add(sec.Title)
		add(sec.ID)
	}
	for _, sc := range doc.scripts {
		add(sc.Name)
		add(sc.RelPath)
	}
	for _, pf := range doc.packageFiles {
		add(pf.Path)
	}
	return fields
}

func scoreSkillQuality(man *SkillManifest, body string, scripts []SkillScriptInfo, sections []SkillSection) (int, []string, map[string]int) {
	score := 0
	warnings := make([]string, 0, 8)
	coverage := map[string]int{
		"sections":          len(sections),
		"scripts":           len(scripts),
		"triggers":          len(manifestTriggers(man)),
		"domains":           len(manifestDomains(man)),
		"role_hints":        len(manifestRoleHints(man)),
		"target_types":      len(manifestTargetTypes(man)),
		"recommended_tools": len(manifestRecommendedTools(man)),
		"required_tools":    len(manifestRequiredTools(man)),
	}

	if strings.TrimSpace(man.Name) != "" {
		score += 10
	} else {
		warnings = append(warnings, "missing manifest name")
	}
	if strings.TrimSpace(man.Description) != "" {
		score += 15
	} else {
		warnings = append(warnings, "missing description")
	}
	if strings.TrimSpace(body) != "" {
		score += 20
	} else {
		warnings = append(warnings, "empty body")
	}
	if len(sections) > 0 {
		score += minInt(len(sections)*8, 24)
	} else {
		warnings = append(warnings, "no markdown sections")
	}
	if len(scripts) > 0 {
		score += minInt(len(scripts)*6, 12)
	}
	if coverage["triggers"] > 0 {
		score += 6
	} else {
		warnings = append(warnings, "no triggers")
	}
	if coverage["domains"] > 0 {
		score += 6
	}
	if coverage["role_hints"] > 0 {
		score += 6
	}
	if coverage["target_types"] > 0 {
		score += 8
	} else {
		warnings = append(warnings, "no target types")
	}
	if coverage["recommended_tools"] > 0 || coverage["required_tools"] > 0 {
		score += 8
	} else {
		warnings = append(warnings, "no tool guidance")
	}

	if score > 100 {
		score = 100
	}
	return score, warnings, coverage
}

func cloneCoverageMap(src map[string]int) map[string]int {
	if len(src) == 0 {
		return nil
	}
	out := make(map[string]int, len(src))
	for k, v := range src {
		out[k] = v
	}
	return out
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func scoreSkillDoc(doc skillPackageDoc, queryTerms, roleTerms []string, availableTools map[string]struct{}, inferredTargets []string) SkillRecommendation {
	rec := SkillRecommendation{
		ID:               doc.summary.ID,
		DirName:          doc.summary.DirName,
		Name:             doc.summary.Name,
		Description:      doc.summary.Description,
		Version:          doc.summary.Version,
		Tags:             append([]string(nil), doc.summary.Tags...),
		Triggers:         append([]string(nil), doc.summary.Triggers...),
		Aliases:          append([]string(nil), doc.summary.Aliases...),
		Domains:          append([]string(nil), doc.summary.Domains...),
		Stages:           append([]string(nil), doc.summary.Stages...),
		TargetTypes:      append([]string(nil), doc.summary.TargetTypes...),
		BundleOf:         append([]string(nil), doc.summary.BundleOf...),
		DependsOn:        append([]string(nil), doc.summary.DependsOn...),
		IsBundle:         doc.summary.IsBundle,
		RecommendedTools: append([]string(nil), doc.summary.RecommendedTools...),
		RequiredTools:    append([]string(nil), doc.summary.RequiredTools...),
		RoleHints:        append([]string(nil), doc.summary.RoleHints...),
		AutoLoadPriority: doc.summary.AutoLoadPriority,
		Score:            doc.summary.AutoLoadPriority,
		QualityScore:     doc.summary.QualityScore,
		QualityWarnings:  append([]string(nil), doc.summary.QualityWarnings...),
	}

	addReason := func(delta int, reason string) {
		if delta <= 0 {
			return
		}
		rec.Score += delta
		if reason != "" {
			rec.Reasons = append(rec.Reasons, reason)
		}
	}

	nameSet := normalizeLowerSet([]string{doc.summary.DirName, doc.summary.Name})
	triggerSet := normalizeLowerSet(doc.summary.Triggers)
	tagSet := normalizeLowerSet(doc.summary.Tags)
	aliasSet := normalizeLowerSet(doc.summary.Aliases)
	domainSet := normalizeLowerSet(doc.summary.Domains)
	stageSet := normalizeLowerSet(doc.summary.Stages)
	targetSet := normalizeLowerSet(doc.summary.TargetTypes)
	bundleSet := normalizeLowerSet(doc.summary.BundleOf)
	dependsSet := normalizeLowerSet(doc.summary.DependsOn)
	roleSet := normalizeLowerSet(doc.summary.RoleHints)
	requiredSet := normalizeLowerSet(doc.summary.RequiredTools)
	recommendedSet := normalizeLowerSet(doc.summary.RecommendedTools)

	for _, term := range queryTerms {
		switch {
		case containsSet(nameSet, term):
			addReason(40, "exact skill name/id hit: "+term)
		case containsSet(aliasSet, term):
			addReason(26, "alias hit: "+term)
		case containsSet(triggerSet, term):
			addReason(22, "trigger hit: "+term)
		case containsSet(targetSet, term):
			addReason(18, "target hit: "+term)
		case containsSet(tagSet, term):
			addReason(16, "tag hit: "+term)
		case containsSet(domainSet, term):
			addReason(14, "domain hit: "+term)
		case containsSet(bundleSet, term):
			addReason(12, "bundle coverage hit: "+term)
		case containsSet(stageSet, term):
			addReason(10, "stage hit: "+term)
		case containsSet(dependsSet, term):
			addReason(8, "dependency hit: "+term)
		case docContains(doc.searchFields, term):
			addReason(8, "content hit: "+term)
		}
	}

	for _, term := range roleTerms {
		switch {
		case containsSet(roleSet, term):
			addReason(18, "role hint hit: "+term)
		case containsSet(targetSet, term):
			addReason(10, "role target hit: "+term)
		case docContains(doc.searchFields, term):
			addReason(6, "role-related content hit: "+term)
		}
	}

	for _, target := range inferredTargets {
		switch {
		case containsSet(targetSet, target):
			addReason(24, "target route hit: "+target)
			rec.MatchedTargets = append(rec.MatchedTargets, target)
		case containsSet(domainSet, target):
			addReason(12, "target-domain hit: "+target)
			rec.MatchedTargets = append(rec.MatchedTargets, target)
		case containsSet(tagSet, target):
			addReason(8, "target-tag hit: "+target)
			rec.MatchedTargets = append(rec.MatchedTargets, target)
		}
	}

	for toolName := range availableTools {
		if containsSet(requiredSet, toolName) {
			addReason(14, "required tool available: "+toolName)
		}
		if containsSet(recommendedSet, toolName) {
			addReason(10, "recommended tool available: "+toolName)
		}
	}

	if rec.IsBundle {
		if n := len(rec.BundleOf); n > 0 {
			addReason(n*2, fmt.Sprintf("bundle breadth: %d", n))
		}
		if n := len(rec.DependsOn); n > 0 {
			addReason(n, fmt.Sprintf("bundle dependencies: %d", n))
		}
	}

	if len(requiredSet) > 0 {
		missing := 0
		for toolName := range requiredSet {
			if _, ok := availableTools[toolName]; !ok {
				missing++
			}
		}
		if missing == 0 && len(availableTools) > 0 {
			addReason(12, "all required tools available")
		} else if missing == len(requiredSet) && len(availableTools) > 0 {
			rec.Score -= 8
		}
	}

	if rec.Score < 0 {
		rec.Score = 0
	}
	rec.MatchedTargets = uniqueStringsKeepOrder(rec.MatchedTargets)
	return rec
}

func attachBundleSignals(recs []SkillRecommendation, bundles []SkillRecommendation) {
	if len(recs) == 0 || len(bundles) == 0 {
		return
	}
	for i := range recs {
		rec := &recs[i]
		for _, bundle := range bundles {
			if !bundleContainsSkill(bundle, rec) {
				continue
			}
			rec.Score += 9
			rec.BundleSources = append(rec.BundleSources, bundle.Name)
			rec.Reasons = append(rec.Reasons, "bundle match: "+bundle.Name)
		}
		rec.BundleSources = uniqueStringsKeepOrder(rec.BundleSources)
	}
}

func bundleContainsSkill(bundle SkillRecommendation, rec *SkillRecommendation) bool {
	if rec == nil || !bundle.IsBundle || len(bundle.BundleOf) == 0 {
		return false
	}
	bundleSet := normalizeLowerSet(bundle.BundleOf)
	for _, candidate := range []string{rec.ID, rec.DirName, rec.Name} {
		if containsSet(bundleSet, candidate) {
			return true
		}
	}
	for _, item := range rec.TargetTypes {
		if containsSet(bundleSet, item) {
			return true
		}
	}
	for _, item := range rec.Tags {
		if containsSet(bundleSet, item) {
			return true
		}
	}
	return false
}

func inferTargetTypes(query, role string, tools []string) []string {
	text := strings.ToLower(strings.TrimSpace(query + " " + role))
	textTerms := normalizeSearchTerms(text)
	textSet := normalizeLowerSet(textTerms)
	toolSet := normalizeLowerSet(tools)

	type scoredTarget struct {
		Name  string
		Score int
	}
	scored := make([]scoredTarget, 0, len(defaultTargetProfiles))
	for _, profile := range defaultTargetProfiles {
		score := 0
		for _, kw := range profile.Keywords {
			if containsTextOrTerm(text, textSet, kw) {
				score += 6
			}
		}
		for _, hint := range profile.RoleHints {
			if containsTextOrTerm(text, textSet, hint) {
				score += 4
			}
		}
		for _, toolName := range profile.ToolHints {
			if _, ok := toolSet[strings.ToLower(toolName)]; ok {
				score += 5
			}
		}
		if score > 0 {
			scored = append(scored, scoredTarget{Name: profile.Name, Score: score})
		}
	}
	sort.Slice(scored, func(i, j int) bool {
		if scored[i].Score != scored[j].Score {
			return scored[i].Score > scored[j].Score
		}
		return scored[i].Name < scored[j].Name
	})

	out := make([]string, 0, len(scored))
	for _, item := range scored {
		out = append(out, item.Name)
		if len(out) >= 3 {
			break
		}
	}
	return out
}

func containsTextOrTerm(text string, terms map[string]struct{}, needle string) bool {
	needle = strings.TrimSpace(strings.ToLower(needle))
	if needle == "" {
		return false
	}
	if strings.Contains(text, needle) {
		return true
	}
	return containsSet(terms, needle)
}

func normalizeSearchTerms(input string) []string {
	if strings.TrimSpace(input) == "" {
		return nil
	}
	parts := strings.FieldsFunc(strings.ToLower(input), func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_')
	})
	return dedupeStrings(parts)
}

func normalizeLowerSet(items []string) map[string]struct{} {
	out := make(map[string]struct{}, len(items))
	for _, item := range items {
		s := strings.TrimSpace(strings.ToLower(item))
		if s == "" {
			continue
		}
		out[s] = struct{}{}
	}
	return out
}

func containsSet(set map[string]struct{}, term string) bool {
	if len(set) == 0 {
		return false
	}
	term = strings.TrimSpace(strings.ToLower(term))
	if term == "" {
		return false
	}
	for item := range set {
		if item == term || strings.Contains(item, term) || strings.Contains(term, item) {
			return true
		}
	}
	return false
}

func docContains(fields []string, term string) bool {
	if len(fields) == 0 {
		return false
	}
	term = strings.TrimSpace(strings.ToLower(term))
	if term == "" {
		return false
	}
	for _, field := range fields {
		if strings.Contains(field, term) {
			return true
		}
	}
	return false
}

func uniqueStringsKeepOrder(items []string) []string {
	if len(items) == 0 {
		return nil
	}
	seen := make(map[string]struct{}, len(items))
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
		seen[key] = struct{}{}
		out = append(out, s)
	}
	return out
}

func cloneMetadata(in map[string]any) map[string]any {
	if len(in) == 0 {
		return nil
	}
	out := make(map[string]any, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

// SkillsRootLabel returns a short label suitable for logs / diagnostics.
func SkillsRootLabel(skillsRoot string) string {
	base := filepath.Base(strings.TrimSpace(skillsRoot))
	if base == "" || base == "." || base == string(filepath.Separator) {
		return skillsRoot
	}
	return base
}
