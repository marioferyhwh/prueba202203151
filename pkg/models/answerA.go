package models

// ComprasPorTDC
type ComprasPorTDC struct {
	Oro  int `json:"oro"`
	Amex int `json:"amex"`
}

//AnswerA
type AnswerA struct {
	Total         float64       `json:"total"`
	ComprasPorTDC ComprasPorTDC `json:"comprasPorTDC"`
	NoCompraron   int           `json:"noCompraron"`
	CompraMasAlta int           `json:"compraMasAlta"`
}
