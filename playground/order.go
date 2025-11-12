package main

import "fmt"

type order struct {
	ProductCode int         `json:"productCode"`
	Quantity    float64     `json:"quantity"`
	Status      orderStatus `json:"status"`
}

func (o order) String() string {
	return fmt.Sprintf("Product code: %v, Quantity: %v, Status: %v\n",
		o.ProductCode, o.Quantity, orderStatusToText(o.Status))
}

func orderStatusToText(status orderStatus) string {
	switch status {
	case none:
		return "none"
	case new:
		return "new"
	case received:
		return "received"
	case reserved:
		return "reserved"
	case filled:
		return "filled"
	default:
		return "unknown status"
	}
}

type orderStatus int

const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)

var orders = []order{}
