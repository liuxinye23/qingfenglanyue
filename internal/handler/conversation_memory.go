package handler

import (
	"fmt"
	"strings"

	"cyberstrike-ai/internal/agent"
	"cyberstrike-ai/internal/database"

	"go.uber.org/zap"
)

func buildConversationMemoryContext(mem *database.ConversationMemory) string {
	if mem == nil {
		return ""
	}

	summary := strings.TrimSpace(mem.SummaryText)
	facts := dedupeMemoryFacts(mem.KeyFacts)
	if summary == "" && len(facts) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("[Conversation memory]\n")
	if summary != "" {
		b.WriteString(summary)
		b.WriteString("\n")
	}
	if len(facts) > 0 {
		b.WriteString("Key facts:\n")
		for _, fact := range facts {
			b.WriteString("- ")
			b.WriteString(fact)
			b.WriteByte('\n')
		}
	}
	return strings.TrimSpace(b.String())
}

func (h *AgentHandler) syncConversationMemory(conversationID, userMessage, assistantMessage string, history []agent.ChatMessage) {
	if h == nil || h.db == nil {
		return
	}

	summary, facts := summarizeConversationMemory(userMessage, assistantMessage, history)
	if summary == "" && len(facts) == 0 {
		return
	}

	if err := h.db.UpsertConversationMemory(conversationID, summary, facts); err != nil {
		h.logger.Warn("sync conversation memory failed", zap.String("conversationId", conversationID), zap.Error(err))
	}
}

func summarizeConversationMemory(userMessage, assistantMessage string, history []agent.ChatMessage) (string, []string) {
	userMessage = strings.TrimSpace(userMessage)
	assistantMessage = strings.TrimSpace(assistantMessage)

	if userMessage == "" && assistantMessage == "" {
		return "", nil
	}

	userShort := compactText(userMessage, 220)
	assistantShort := compactText(assistantMessage, 320)
	summary := ""
	switch {
	case userShort != "" && assistantShort != "":
		summary = fmt.Sprintf("Latest request: %s\nLatest answer: %s", userShort, assistantShort)
	case userShort != "":
		summary = "Latest request: " + userShort
	case assistantShort != "":
		summary = "Latest answer: " + assistantShort
	}

	facts := make([]string, 0, 6)
	if userShort != "" {
		facts = append(facts, "User focus: "+userShort)
	}
	if assistantShort != "" {
		facts = append(facts, "Latest outcome: "+assistantShort)
	}

	if len(history) > 0 {
		lastAssistant := ""
		lastUser := ""
		for i := len(history) - 1; i >= 0; i-- {
			msg := history[i]
			role := strings.ToLower(strings.TrimSpace(msg.Role))
			switch {
			case lastAssistant == "" && role == "assistant" && strings.TrimSpace(msg.Content) != "":
				lastAssistant = compactText(msg.Content, 200)
			case lastUser == "" && role == "user" && strings.TrimSpace(msg.Content) != "":
				lastUser = compactText(msg.Content, 160)
			}
			if lastAssistant != "" && lastUser != "" {
				break
			}
		}
		if lastUser != "" {
			facts = append(facts, "Previous user context: "+lastUser)
		}
		if lastAssistant != "" {
			facts = append(facts, "Previous assistant context: "+lastAssistant)
		}
	}

	return summary, dedupeMemoryFacts(facts)
}

func dedupeMemoryFacts(items []string) []string {
	if len(items) == 0 {
		return nil
	}
	seen := make(map[string]struct{}, len(items))
	out := make([]string, 0, len(items))
	for _, item := range items {
		item = compactText(item, 220)
		if item == "" {
			continue
		}
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}
	return out
}

func compactText(s string, maxLen int) string {
	s = strings.TrimSpace(strings.Join(strings.Fields(s), " "))
	if s == "" || maxLen <= 0 {
		return s
	}
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	return string(runes[:maxLen]) + "..."
}
