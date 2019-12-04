package models


type Resp map[string]string

type FullResp struct {
	Status string `json:"status"`
	Depot []Resp  `json:"depot,omitempty"`
	Routes []Resp `json:"routes,omitempty"`
}

type Discount struct {
	Status string `json:"status"`
	Busclassid string `json:"busClassId"`
	Busorderid string  `json:"busOrderId"`
	Discountid string `json:"discountId"`
	Discountname string `json:"discountName"`
	Discountamount string `json:"discountAmount"`
}