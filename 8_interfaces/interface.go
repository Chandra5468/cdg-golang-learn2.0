package main

import "fmt"

// Interfaces are like contract. Whatever struct implements this contract must implement method mentioned in interface

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p *payment) makePayment(amount float32) {
	// razorPayPaymentGw := &razorPay{}

	// razorPayPaymentGw.pay(amount)

	// stripePaymentGW := &stripePay{}
	// stripePaymentGW.pay(amount)

	// WE are violating open-close principle (Open for extension, close for modification) by modifying makePayment instead of extending it.
	// To help with this open-close solid principle. We are using interfaces.
}

type razorPay struct {
}

func (r *razorPay) pay(amount float32) {
	// Logic to make payment
	fmt.Println("Making payment using razor pay ", amount)
}

type stripePay struct {
}

func (s *stripePay) pay(amount float32) {
	fmt.Println("Making payment using stripe pay ", amount)
}

type fakePayment struct {
}

func (f *fakePayment) pay(amount float32) {
	fmt.Println("Making payment from fake gateway for testing ", amount)
}
func main() {
	// Very Heavily used concept to organize and scale a project

	// newPayment := &payment{}
	// newPayment.makePayment(100)

	// Below interface implementation is good for SOLID- open/close principle. As enhancement can be done instead of modification.

	// razorPayGw := &razorPay{}

	// newPayment := payment{
	// 	gateway: razorPayGw,
	// }

	// newPayment.gateway.pay(110)

	fakePaymentGW := &fakePayment{}

	newPayment := payment{
		gateway: fakePaymentGW,
	}

	newPayment.gateway.pay(504)
}
