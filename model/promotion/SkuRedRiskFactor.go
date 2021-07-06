package promotion

// SkuRedRiskFactor 结构体
type SkuRedRiskFactor struct {
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// sku红线价格
	AmountAt int64 `json:"amount_at,omitempty" xml:"amount_at,omitempty"`
}
