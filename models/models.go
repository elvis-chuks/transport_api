package models

type Resp map[string]string

type Respy map[string]interface{}

type TripResp struct {
	Queueid       string   `json:"busqueueid,omitempty"`
	Classid       string   `json:"busclassid,omitempty"`
	Orderid       string   `json:"busorderid,omitempty"`
	Arrangementid string   `json:"busseatarrangementid,omitempty"`
	NumberOfSeat  string   `json:"numberofseat,omitempty"`
	BookedSeats   []string `json:"bookedseats,omitempty"`
	FreeSeats     string   `json:"freeseats,omitempty"`
	Price         string   `json:"price,omitempty"`
}

type FullResp struct {
	Status  string  `json:"status"`
	Depot   []Resp  `json:"depot,omitempty"`
	Routes  []Resp  `json:"routes,omitempty"`
	Buses   []Resp  `json:"buses,omitempty"`
	BusSeat []Resp  `json:"busSeat,omitempty"`
	Seats   []Resp  `json:"seats,omitempty"`
	Prices  []Resp  `json:"prices,omitempty"`
	Trips   []Respy `json:"trips,omitempty"`
}

type Discount struct {
	Status         string `json:"status"`
	Busclassid     string `json:"busClassId"`
	Busorderid     string `json:"busOrderId"`
	Discountid     string `json:"discountId"`
	Discountname   string `json:"discountName"`
	Discountamount string `json:"discountAmount"`
}
