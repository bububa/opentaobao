package promotion

// SellerGlobalDiscount 结构体
type SellerGlobalDiscount struct {
	// 折扣1折100，9折900
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}
