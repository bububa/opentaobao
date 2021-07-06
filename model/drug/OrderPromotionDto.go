package drug

// OrderPromotionDto 结构体
type OrderPromotionDto struct {
	// 优惠名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 优惠类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 优惠金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}
