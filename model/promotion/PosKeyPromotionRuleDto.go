package promotion

// PosKeyPromotionRuleDto 
type PosKeyPromotionRuleDto struct {
    // 优惠规则
    DetailList   []PosKeyPromotionRuleDetailDto `json:"detail_list,omitempty" xml:"detail_list>pos_key_promotion_rule_detail_dto,omitempty"`
    // 名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // postkey
    PosKey   int64 `json:"pos_key,omitempty" xml:"pos_key,omitempty"`
}
