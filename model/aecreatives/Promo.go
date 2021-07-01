package aecreatives

// Promo 结构体
type Promo struct {
	// 主题活动的商品数量
	ProductNum int64 `json:"product_num,omitempty" xml:"product_num,omitempty"`
	// 主题活动描述
	PromoDesc string `json:"promo_desc,omitempty" xml:"promo_desc,omitempty"`
	// 主题活动名称
	PromoName string `json:"promo_name,omitempty" xml:"promo_name,omitempty"`
}
