package promotion

// PosKeyPromotionRuleDetailDto 
type PosKeyPromotionRuleDetailDto struct {

    // 扩展属性
    ExtMap   string `json:"ext_map,omitempty"`

    // 数量
    Num   int64 `json:"num,omitempty"`

    // 1,9,12,15四个等级
    UnitPoint   string `json:"unit_point,omitempty"`

    // 价格
    UnitPrice   int64 `json:"unit_price,omitempty"`

}
