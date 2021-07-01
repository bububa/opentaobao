package promotion

// RedRiskUpdateFactor 结构体
type RedRiskUpdateFactor struct {
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 需要删除的sku红线价格
	RemoveSkuIds []int64 `json:"remove_sku_ids,omitempty" xml:"remove_sku_ids>int64,omitempty"`
	// 商品红线价格
	AmountAt int64 `json:"amount_at,omitempty" xml:"amount_at,omitempty"`
	// 新增sku级别红线价格
	SkuRiskFactors []SkuRedRiskFactor `json:"sku_risk_factors,omitempty" xml:"sku_risk_factors>sku_red_risk_factor,omitempty"`
	// 风险等级设置
	RiskLevels []RiskLevelParam `json:"risk_levels,omitempty" xml:"risk_levels>risk_level_param,omitempty"`
}
