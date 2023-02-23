package wdk

// UniqueDiscountCodeRequest 结构体
type UniqueDiscountCodeRequest struct {
	// 过期时间。不传默认为当前时间+90天。最大支持有效期为90天后。
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// 门店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 折扣率 范围(0,100) 比如78折传78，8折传80。对应 discountType=7时必传
	DiscountRate int64 `json:"discount_rate,omitempty" xml:"discount_rate,omitempty"`
	// 一口价。单位分。对应 discountType=8 时必传
	DiscountPrice int64 `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
	// 折扣码类型, 7为折扣率码，8为一口价码
	DiscountType int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
}
