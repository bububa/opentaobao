package promotion

// RedRiskUpdateFactor 
/* model for simplify = false
type RedRiskUpdateFactor struct {

    // 商品id
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 需要删除的sku红线价格
    
    RemoveSkuIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"remove_sku_ids,omitempty"`
    

    // 商品红线价格
    
    AmountAt   int64 `json:"amount_at,omitempty"`
    

    // 新增sku级别红线价格
    
    SkuRiskFactors  struct {
        SkuRedRiskFactor  []SkuRedRiskFactor `json:"sku_red_risk_factor,omitempty"`
    } `json:"sku_risk_factors,omitempty"`
    

    // 风险等级设置
    
    RiskLevels  struct {
        RiskLevelParam  []RiskLevelParam `json:"risk_level_param,omitempty"`
    } `json:"risk_levels,omitempty"`
    

}
*/

// RedRiskUpdateFactor 
type RedRiskUpdateFactor struct {

    // 商品id
    ItemId   int64 `json:"item_id,omitempty"`

    // 需要删除的sku红线价格
    RemoveSkuIds   []int64 `json:"remove_sku_ids,omitempty"`

    // 商品红线价格
    AmountAt   int64 `json:"amount_at,omitempty"`

    // 新增sku级别红线价格
    SkuRiskFactors   []SkuRedRiskFactor `json:"sku_risk_factors,omitempty"`

    // 风险等级设置
    RiskLevels   []RiskLevelParam `json:"risk_levels,omitempty"`

}
