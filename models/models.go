package models


type Resp map[string]string

type FullResp struct {
	Status string `json:"status"`
	Depot []Resp  `json:"depot,omitempty"`
	Routes []Resp `json:"routes,omitempty"`
}