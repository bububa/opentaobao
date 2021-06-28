package promotion

// PosKeyPromotionRuleDto 
/* model for simplify = false
type PosKeyPromotionRuleDto struct {

    // 优惠规则
    
    DetailList  struct {
        PosKeyPromotionRuleDetailDto  []PosKeyPromotionRuleDetailDto `json:"pos_key_promotion_rule_detail_dto,omitempty"`
    } `json:"detail_list,omitempty"`
    

    // 名称
    
    Name   string `json:"name,omitempty"`
    

    // postkey
    
    PosKey   int64 `json:"pos_key,omitempty"`
    

}
*/

// PosKeyPromotionRuleDto 
type PosKeyPromotionRuleDto struct {

    // 优惠规则
    DetailList   []PosKeyPromotionRuleDetailDto `json:"detail_list,omitempty"`

    // 名称
    Name   string `json:"name,omitempty"`

    // postkey
    PosKey   int64 `json:"pos_key,omitempty"`

}
