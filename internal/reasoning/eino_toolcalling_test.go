package reasoning

import (
	"testing"

	"cyberstrike-ai/internal/config"

	einoopenai "github.com/cloudwego/eino-ext/components/model/openai"
)

func TestApplyToEinoToolCallingChatModelConfig_StripsDeepseekReasoningFields(t *testing.T) {
	cfg := &einoopenai.ChatModelConfig{}
	oa := &config.OpenAIConfig{
		BaseURL: "https://api.deepseek.com/v1",
		Model:   "deepseek-chat",
		Reasoning: config.OpenAIReasoningConfig{
			Mode: "auto",
			ExtraRequestFields: map[string]any{
				"thinking":         map[string]any{"type": "enabled"},
				"reasoning_effort": "high",
				"foo":              "bar",
			},
		},
	}

	ApplyToEinoToolCallingChatModelConfig(cfg, oa, nil)

	if cfg.ExtraFields == nil {
		t.Fatalf("expected extra fields to be preserved")
	}
	if got := cfg.ExtraFields["foo"]; got != "bar" {
		t.Fatalf("expected non-reasoning extra field to survive, got %#v", got)
	}
	if _, exists := cfg.ExtraFields["thinking"]; exists {
		t.Fatalf("thinking should be stripped for deepseek tool-calling")
	}
	if _, exists := cfg.ExtraFields["reasoning_effort"]; exists {
		t.Fatalf("reasoning_effort should be stripped for deepseek tool-calling")
	}
	if cfg.ReasoningEffort != "" {
		t.Fatalf("reasoning effort should stay empty, got %q", cfg.ReasoningEffort)
	}
}

func TestApplyToEinoToolCallingChatModelConfig_PreservesOpenAIReasoning(t *testing.T) {
	cfg := &einoopenai.ChatModelConfig{}
	oa := &config.OpenAIConfig{
		BaseURL: "https://api.openai.com/v1",
		Model:   "gpt-4o",
		Reasoning: config.OpenAIReasoningConfig{
			Mode:   "on",
			Effort: "high",
		},
	}

	ApplyToEinoToolCallingChatModelConfig(cfg, oa, nil)

	if cfg.ReasoningEffort != einoopenai.ReasoningEffortLevelHigh {
		t.Fatalf("expected high reasoning effort, got %q", cfg.ReasoningEffort)
	}
}
