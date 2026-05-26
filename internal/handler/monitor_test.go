package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"cyberstrike-ai/internal/config"
	"cyberstrike-ai/internal/mcp"
	"cyberstrike-ai/internal/security"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestMonitorHandler_GetToolsHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	logger := zap.NewNop()
	mcpServer := mcp.NewServer(logger)
	secCfg := &config.SecurityConfig{
		Tools: []config.ToolConfig{
			{
				Name:    "internal-tool",
				Command: "internal:query_execution_result",
				Enabled: true,
			},
			{
				Name:    "missing-binary",
				Command: "definitely-not-installed-monitor-test-binary",
				Enabled: true,
			},
		},
	}
	setSecurityConfigWarnings(secCfg, []string{"tool configuration directory not found: /tmp/tools"})
	executor := security.NewExecutor(secCfg, mcpServer, logger)

	mcpServer.RegisterTool(mcp.Tool{
		Name:             "internal-tool",
		Description:      "internal configured tool",
		ShortDescription: "internal configured tool",
	}, func(ctx context.Context, args map[string]interface{}) (*mcp.ToolResult, error) {
		return &mcp.ToolResult{}, nil
	})
	mcpServer.RegisterTool(mcp.Tool{
		Name:             "ad-hoc-tool",
		Description:      "ad hoc direct MCP tool",
		ShortDescription: "ad hoc direct MCP tool",
		InputSchema:      map[string]interface{}{"type": "object"},
	}, func(ctx context.Context, args map[string]interface{}) (*mcp.ToolResult, error) {
		return &mcp.ToolResult{}, nil
	})

	handler := NewMonitorHandler(mcpServer, executor, nil, logger)
	router := gin.New()
	router.GET("/api/monitor/tools-health", handler.GetToolsHealth)

	req := httptest.NewRequest("GET", "/api/monitor/tools-health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp ToolsHealthResponse
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(resp.Configured) != 2 {
		t.Fatalf("expected 2 configured tools, got %d", len(resp.Configured))
	}
	if len(resp.DirectMCP) != 1 || resp.DirectMCP[0].Name != "ad-hoc-tool" {
		t.Fatalf("unexpected direct MCP tools: %#v", resp.DirectMCP)
	}
	if resp.Summary["configured_total"] != 2 {
		t.Fatalf("unexpected summary: %#v", resp.Summary)
	}
	if securityConfigSupportsWarnings() {
		if len(resp.Warnings) != 1 || resp.Warnings[0] != "tool configuration directory not found: /tmp/tools" {
			t.Fatalf("unexpected warnings: %#v", resp.Warnings)
		}
	} else if len(resp.Warnings) != 0 {
		t.Fatalf("expected no warnings on configs without ToolConfigWarnings, got %#v", resp.Warnings)
	}
	if resp.StatusBuckets["ok"] != 1 {
		t.Fatalf("expected ok status bucket, got %#v", resp.StatusBuckets)
	}
	if resp.StatusBuckets["missing_binary"] != 1 {
		t.Fatalf("expected missing_binary status bucket, got %#v", resp.StatusBuckets)
	}
}

func setSecurityConfigWarnings(cfg *config.SecurityConfig, warnings []string) {
	if cfg == nil {
		return
	}
	val := reflect.ValueOf(cfg).Elem()
	field := val.FieldByName("ToolConfigWarnings")
	if !field.IsValid() || !field.CanSet() || field.Kind() != reflect.Slice {
		return
	}
	field.Set(reflect.ValueOf(warnings))
}

func securityConfigSupportsWarnings() bool {
	_, ok := reflect.TypeOf(config.SecurityConfig{}).FieldByName("ToolConfigWarnings")
	return ok
}
