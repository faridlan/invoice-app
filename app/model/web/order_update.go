package web

type OrderUpdate struct {
	Id        string `json:"id,omitempty"`
	OrderDate int64  `json:"order_date,omitempty"`
	CusName   string `json:"cus_name,omitempty"`
	Total     int    `json:"total,omitempty"`
	Dp        int    `json:"dp,omitempty"`
	Pay       int    `json:"pay,omitempty"`
	RestOfPay int    `json:"rest_of_pay,omitempty"`
}
