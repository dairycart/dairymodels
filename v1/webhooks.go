package models

import (
	"time"
)

// Webhook represents a Dairycart webhook
type Webhook struct {
	ArchivedOn  NullTime  `json:"archived_on"`  // archived_on
	EventType   string    `json:"event_type"`   // event_type
	ID          uint64    `json:"id"`           // id
	ContentType string    `json:"content_type"` // content_type
	UpdatedOn   NullTime  `json:"updated_on"`   // updated_on
	CreatedOn   time.Time `json:"created_on"`   // created_on
	URL         string    `json:"url"`          // url
}

type WebhookListResponse struct {
	ListResponse
	Webhooks []Webhook `json: "webhooks"`
}
