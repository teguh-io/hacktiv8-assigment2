package params

import "time"

type Order struct {
	ID           int        `json:"id,omitempty"`
	CustomerName *string    `json:"customer_name"`
	Items        []Item     `json:"items,omitempty"`
	OrderedAt    *time.Time `json:"ordered_at,omitempty"`
}

type OrderResponse struct {
	Status         int     `json:"status"`
	Message        *string `json:"message,omitempty"`
	AdditionalInfo *string `json:"additional_info,omitempty"`
	Payload        Order   `json:"payload,omitempty"`
}
