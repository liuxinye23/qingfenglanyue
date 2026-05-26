package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ConversationMemory stores compact long-lived context for one conversation.
type ConversationMemory struct {
	ConversationID string    `json:"conversation_id"`
	SummaryText    string    `json:"summary_text"`
	KeyFacts       []string  `json:"key_facts,omitempty"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GetConversationMemory loads persisted memory for a conversation.
func (db *DB) GetConversationMemory(conversationID string) (*ConversationMemory, error) {
	if strings.TrimSpace(conversationID) == "" {
		return nil, fmt.Errorf("conversationID is empty")
	}

	var summary sql.NullString
	var keyFactsJSON sql.NullString
	var updatedAt sql.NullTime
	mem := &ConversationMemory{ConversationID: conversationID}
	err := db.QueryRow(
		`SELECT summary_text, key_facts_json, updated_at
		 FROM conversation_memories
		 WHERE conversation_id = ?`,
		conversationID,
	).Scan(&summary, &keyFactsJSON, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("load conversation memory: %w", err)
	}

	if summary.Valid {
		mem.SummaryText = strings.TrimSpace(summary.String)
	}
	if keyFactsJSON.Valid && strings.TrimSpace(keyFactsJSON.String) != "" {
		if err := json.Unmarshal([]byte(keyFactsJSON.String), &mem.KeyFacts); err != nil {
			return nil, fmt.Errorf("decode conversation memory key facts: %w", err)
		}
	}
	if updatedAt.Valid {
		mem.UpdatedAt = updatedAt.Time
	}
	return mem, nil
}

// UpsertConversationMemory stores or updates persisted memory for a conversation.
func (db *DB) UpsertConversationMemory(conversationID, summaryText string, keyFacts []string) error {
	conversationID = strings.TrimSpace(conversationID)
	if conversationID == "" {
		return fmt.Errorf("conversationID is empty")
	}

	summaryText = strings.TrimSpace(summaryText)
	keyFacts = uniqueNonEmptyStrings(keyFacts)
	now := time.Now()

	var keyFactsJSON string
	if len(keyFacts) > 0 {
		b, err := json.Marshal(keyFacts)
		if err != nil {
			return fmt.Errorf("encode conversation memory key facts: %w", err)
		}
		keyFactsJSON = string(b)
	}

	_, err := db.Exec(
		`INSERT INTO conversation_memories (conversation_id, summary_text, key_facts_json, updated_at)
		 VALUES (?, ?, ?, ?)
		 ON CONFLICT(conversation_id) DO UPDATE SET
		   summary_text = excluded.summary_text,
		   key_facts_json = excluded.key_facts_json,
		   updated_at = excluded.updated_at`,
		conversationID,
		summaryText,
		keyFactsJSON,
		now,
	)
	if err != nil {
		return fmt.Errorf("upsert conversation memory: %w", err)
	}
	return nil
}

func uniqueNonEmptyStrings(items []string) []string {
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
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}
	return out
}
