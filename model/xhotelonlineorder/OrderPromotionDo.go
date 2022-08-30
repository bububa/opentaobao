package xhotelonlineorder

// OrderPromotionDo 结构体
type OrderPromotionDo struct {
	// OrderPromotionDetail
	CurrentOrderPromotion *OrderPromotionDetail `json:"current_order_promotion,omitempty" xml:"current_order_promotion,omitempty"`
}
