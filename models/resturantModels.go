package models

import "time"

// Customers model
type Customer struct {
	CustomerID  int    `json:"customerID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

// Orders model
type Order struct {
	OrderID       int       `json:"orderID"`
	CustomerID    int       `json:"customerID"`
	OrderDateTime time.Time `json:"orderDateTime"`
	TotalPrice    int64     `json:"totalPrice"`
}

// OrderItems model
type OrderItem struct {
	OrderItemID int   `json:"orderItemID"`
	OrderID     int   `json:"orderID"`
	ItemID      int   `json:"itemID"`
	Quantity    int   `json:"quantity"`
	Subtotal    int64 `json:"subtotal"`
}

// Staff model
type Staff struct {
	StaffID     int    `json:"staffID" `
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

// Payment model
type Payment struct {
	PaymentID       int       `json:"paymentID" gorm:"column:payment_id"`
	OrderID         int       `json:"orderID" gorm:"column:order_id"`
	PaymentDateTime time.Time `json:"paymentDateTime" gorm:"column:payment_date_time"`
	PaymentMethod   string    `json:"paymentMethod" gorm:"column:payment_method"`
}
