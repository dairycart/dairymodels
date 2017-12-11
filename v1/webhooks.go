package models

// Webhook represents a Dairycart webhook
type Webhook struct {
	ID          uint64     `json:"id"`           // id
	URL         string     `json:"url"`          // url
	EventType   string     `json:"event_type"`   // event_type
	ContentType string     `json:"content_type"` // content_type
	CreatedOn   *Dairytime `json:"created_on"`   // created_on
	UpdatedOn   *Dairytime `json:"updated_on"`   // updated_on
	ArchivedOn  *Dairytime `json:"archived_on"`  // archived_on
}

type WebhookCreationInput struct {
	ID          uint64 `json:"id,omitempty"`           // id
	URL         string `json:"url,omitempty"`          // url
	EventType   string `json:"event_type,omitempty"`   // event_type
	ContentType string `json:"content_type,omitempty"` // content_type

}

type WebhookListResponse struct {
	ListResponse
	Webhooks []Webhook `json:"webhooks"`
}
