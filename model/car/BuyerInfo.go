package car

// BuyerInfo 结构体
type BuyerInfo struct {
	// buyerEmail
	BuyerEmail string `json:"buyer_email,omitempty" xml:"buyer_email,omitempty"`
	// buyerPhone
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
}
