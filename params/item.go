package params

type Item struct {
	ItemCode    *string `json:"item_code"`
	Description *string `json:"description"`
	Quantity    *int    `json:"quantity"`
	OrderID     *int    `json:"OrderID,omitempty"`
}
