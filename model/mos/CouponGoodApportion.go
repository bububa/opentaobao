package mos

// CouponGoodApportion 结构体
type CouponGoodApportion struct {
	// 商品行号
	LineNo string `json:"line_no,omitempty" xml:"line_no,omitempty"`
	// 商品sku 如果有的话
	Sku string `json:"sku,omitempty" xml:"sku,omitempty"`
	// 商品在该券里分摊的金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
}
