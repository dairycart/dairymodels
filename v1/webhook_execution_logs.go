package models

import (
	"time"
)

// WebhookExecutionLog represents a Dairycart webhook execution log
type WebhookExecutionLog struct {
	ID         uint64    `json:"id"`          // id
	Succeeded  bool      `json:"succeeded"`   // succeeded
	StatusCode int       `json:"status_code"` // status_code
	ExecutedOn time.Time `json:"executed_on"` // executed_on
	WebhookID  uint64    `json:"webhook_id"`  // webhook_id
}

type WebhookExecutionLogListResponse struct {
	ListResponse
	WebhookExecutionLogs []WebhookExecutionLog `json:"webhook_execution_logs"`
}
