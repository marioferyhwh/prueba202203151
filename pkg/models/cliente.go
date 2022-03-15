package models

import "time"

// Client .
type Client struct {
	ClientID uint64    `json"clientId"`
	Phone    string    `json:"phone,omitempty"`
	Nombre   string    `json:"nombre,omitempty"`
	Compro   bool      `json:"compro,omitempty"`
	Tdc      string    `json:"tdc,omitempty"`
	Monto    float64   `json:"monto,omitempty"`
	Date     time.Time `json:"date,omitempty"`
}
