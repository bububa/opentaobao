package promotion

// PosKeyPromotionRuleDetailDTO 
type PosKeyPromotionRuleDetailDTO struct {
    // 扩展属性
    ExtMap   string `json:"ext_map,omitempty" xml:"ext_map,omitempty"`
    // 数量
    Num   int64 `json:"num,omitempty" xml:"num,omitempty"`
    // 1,9,12,15四个等级
    UnitPoint   string `json:"unit_point,omitempty" xml:"unit_point,omitempty"`
    // 价格
    UnitPrice   int64 `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
}
