package domain

type DetailOrder struct {
	Id          string `json:"id,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	UnitPrice   int    `json:"unit_price,omitempty"`
}
