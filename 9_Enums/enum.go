package main

import "fmt"

// Enumerated types : to name a list of things separately, one by one

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

type OrderStatusString string

const (
	rcvd OrderStatusString = "received"
	cnfd                   = "Confirmed"
	prpd                   = "Prepared"
	dlvd                   = "Delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Chaning order status to ", status)
}

func changeOrderStatusString(status OrderStatusString) {
	fmt.Println("changing strinfy order status to ", status)
}
func main() {
	changeOrderStatus(Received)

	changeOrderStatusString(cnfd)
}
