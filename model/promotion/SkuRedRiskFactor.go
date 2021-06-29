package promotion

// SkuRedRiskFactor 
type SkuRedRiskFactor struct {
    // sku红线价格
    AmountAt   int64 `json:"amount_at,omitempty" xml:"amount_at,omitempty"`
    // skuId
    SkuId   string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
