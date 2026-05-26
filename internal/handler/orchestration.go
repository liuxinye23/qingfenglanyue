package handler

import (
	"strings"

	"cyberstrike-ai/internal/config"
)

func decideEffectiveOrchestration(req *ChatRequest, finalMessage string, roleTools []string) string {
	if req != nil {
		if requested := strings.TrimSpace(req.Orchestration); requested != "" {
			return config.NormalizeMultiAgentOrchestration(requested)
		}
	}

	lower := strings.ToLower(finalMessage)
	toolCount := len(roleTools)

	switch {
	case strings.Contains(lower, "delegate"), strings.Contains(lower, "sub-agent"), strings.Contains(lower, "sub agent"), strings.Contains(lower, "parallel"), toolCount >= 8:
		return "supervisor"
	case strings.Contains(lower, "step by step"), strings.Contains(lower, "plan"), strings.Contains(lower, "first "), strings.Contains(lower, "然后"), strings.Contains(lower, "步骤"), strings.Contains(lower, "分阶段"):
		return "plan_execute"
	default:
		return "deep"
	}
}
