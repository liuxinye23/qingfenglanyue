package multiagent

import (
	"context"
	"strings"

	"cyberstrike-ai/internal/skillpackage"

	"github.com/cloudwego/eino/components/tool"
	"go.uber.org/zap"
)

func injectSkillShortlistInstruction(
	ctx context.Context,
	instruction string,
	skillsRoot string,
	userMessage string,
	roleTools []string,
	tools []tool.BaseTool,
	logger *zap.Logger,
) string {
	root := strings.TrimSpace(skillsRoot)
	if root == "" {
		return strings.TrimSpace(instruction)
	}

	availableTools := make([]string, 0, len(roleTools)+len(tools))
	availableTools = append(availableTools, roleTools...)
	availableTools = append(availableTools, collectToolNames(ctx, tools)...)

	shortlistText, recs, err := skillpackage.BuildSkillShortlistInstruction(root, skillpackage.SkillRecommendOptions{
		Query:          userMessage,
		AvailableTools: availableTools,
		Limit:          3,
	})
	if err != nil {
		if logger != nil {
			logger.Warn("skill shortlist build failed", zap.String("skills_root", skillpackage.SkillsRootLabel(root)), zap.Error(err))
		}
		return strings.TrimSpace(instruction)
	}
	if strings.TrimSpace(shortlistText) == "" {
		return strings.TrimSpace(instruction)
	}
	if logger != nil {
		names := make([]string, 0, len(recs))
		for _, rec := range recs {
			names = append(names, rec.Name)
		}
		logger.Info("skill shortlist injected",
			zap.String("skills_root", skillpackage.SkillsRootLabel(root)),
			zap.Strings("skills", names),
			zap.Int("count", len(recs)),
		)
	}

	var sb strings.Builder
	sb.WriteString(shortlistText)
	if s := strings.TrimSpace(instruction); s != "" {
		sb.WriteString("\n\n")
		sb.WriteString(s)
	}
	return sb.String()
}
