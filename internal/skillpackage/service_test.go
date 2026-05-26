package skillpackage

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRecommendSkillsDetailedPrefersBundleAndMembers(t *testing.T) {
	root := t.TempDir()

	writeSkill(t, root, "auth-session-review", `---
name: auth-session-review
description: Review sessions.
target_types:
  - web
  - api
tags:
  - auth
  - session
triggers:
  - login
  - cookie
recommended_tools:
  - http-framework-test
autoload_priority: 10
---

# Auth Session Review
`)

	writeSkill(t, root, "token-lifecycle-review", `---
name: token-lifecycle-review
description: Review token lifecycle.
target_types:
  - api
  - web
tags:
  - token
  - auth
triggers:
  - refresh token
  - bearer token
recommended_tools:
  - jwt-analyzer
autoload_priority: 9
---

# Token Lifecycle Review
`)

	writeSkill(t, root, "web-auth-bundle", `---
name: web-auth-bundle
description: Bundle for web auth review.
target_types:
  - web
  - api
bundle_of:
  - auth-session-review
  - token-lifecycle-review
triggers:
  - login flow
  - session cookie
recommended_tools:
  - http-framework-test
  - jwt-analyzer
autoload_priority: 20
---

# Web Auth Bundle
`)

	result, err := RecommendSkillsDetailed(root, SkillRecommendOptions{
		Query:          "review login flow, session cookie and bearer token refresh issues",
		AvailableTools: []string{"http-framework-test", "jwt-analyzer"},
		Limit:          5,
	})
	if err != nil {
		t.Fatalf("RecommendSkillsDetailed returned error: %v", err)
	}
	if got := result.PrimaryTarget; got != "web" && got != "api" {
		t.Fatalf("unexpected primary target: %q", got)
	}
	if len(result.Bundles) == 0 || result.Bundles[0].DirName != "web-auth-bundle" {
		t.Fatalf("expected web-auth-bundle as top bundle, got %#v", result.Bundles)
	}
	if len(result.Recommendations) < 2 {
		t.Fatalf("expected member recommendations, got %#v", result.Recommendations)
	}
	if result.Recommendations[0].DirName != "auth-session-review" {
		t.Fatalf("expected auth-session-review first, got %q", result.Recommendations[0].DirName)
	}
	if !containsString(result.Recommendations[0].BundleSources, "web-auth-bundle") {
		t.Fatalf("expected member to inherit bundle source, got %#v", result.Recommendations[0].BundleSources)
	}
}

func TestBuildSkillShortlistInstructionIncludesRouteAndBundle(t *testing.T) {
	root := t.TempDir()

	writeSkill(t, root, "ctf-skills", `---
name: ctf-skills
description: Generic ctf router.
target_types:
  - ctf
triggers:
  - ctf
  - challenge
autoload_priority: 15
---

# CTF Skills
`)

	writeSkill(t, root, "elf-pwn-triage", `---
name: elf-pwn-triage
description: ELF pwn triage.
target_types:
  - ctf
  - binary
  - reverse
tags:
  - pwn
  - elf
triggers:
  - elf
  - nc
recommended_tools:
  - gdb
autoload_priority: 18
---

# ELF Pwn Triage
`)

	writeSkill(t, root, "ctf-binary-bundle", `---
name: ctf-binary-bundle
description: Binary ctf bundle.
target_types:
  - ctf
  - binary
bundle_of:
  - ctf-skills
  - elf-pwn-triage
triggers:
  - binary challenge
  - elf
recommended_tools:
  - gdb
autoload_priority: 25
---

# CTF Binary Bundle
`)

	text, recs, err := BuildSkillShortlistInstruction(root, SkillRecommendOptions{
		Query:          "ctf elf binary over nc",
		AvailableTools: []string{"gdb"},
		Limit:          3,
	})
	if err != nil {
		t.Fatalf("BuildSkillShortlistInstruction returned error: %v", err)
	}
	if len(recs) == 0 {
		t.Fatal("expected shortlist recommendations")
	}
	if !strings.Contains(text, "primary_target=") {
		t.Fatalf("expected route info in shortlist: %s", text)
	}
	if !strings.Contains(text, "bundles: ctf-binary-bundle") {
		t.Fatalf("expected bundle info in shortlist: %s", text)
	}
	if !strings.Contains(text, "bundled_by=ctf-binary-bundle") {
		t.Fatalf("expected bundled_by marker in shortlist: %s", text)
	}
}

func TestLoadSkillIncludesQualityScore(t *testing.T) {
	root := t.TempDir()
	writeSkill(t, root, "quality-check-skill", `---
name: quality-check-skill
description: Quality signal test.
target_types:
  - web
triggers:
  - auth
recommended_tools:
  - http-framework-test
role_hints:
  - web
---

# Quality Check

## Workflow
Run the documented workflow.
`)

	v, err := LoadSkill(root, "quality-check-skill", LoadOptions{Depth: "full"})
	if err != nil {
		t.Fatalf("LoadSkill returned error: %v", err)
	}
	if v.QualityScore <= 0 {
		t.Fatalf("expected positive quality score, got %d", v.QualityScore)
	}
	if v.Coverage["sections"] == 0 {
		t.Fatalf("expected section coverage, got %#v", v.Coverage)
	}
}

func writeSkill(t *testing.T, root, name, content string) {
	t.Helper()
	dir := filepath.Join(root, name)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir skill %s: %v", name, err)
	}
	if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte(content), 0o644); err != nil {
		t.Fatalf("write skill %s: %v", name, err)
	}
}

func containsString(items []string, want string) bool {
	for _, item := range items {
		if item == want {
			return true
		}
	}
	return false
}
