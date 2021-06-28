package promotion

// Range 
/* model for simplify = false
type Range struct {

    // 营销范围参与者id。即MarketingRangeDO中的participateId。
    
    ParticipateId   int64 `json:"participate_id,omitempty"`
    

    // 营销范围参与者类型。MarketingRangeDO中的participateType。(1:商品;2:店铺;3:seller;4:sku;5:category;6:shopCategory)
    
    ParticipateType   int64 `json:"participate_type,omitempty"`
    

}
*/

// Range 
type Range struct {

    // 营销范围参与者id。即MarketingRangeDO中的participateId。
    ParticipateId   int64 `json:"participate_id,omitempty"`

    // 营销范围参与者类型。MarketingRangeDO中的participateType。(1:商品;2:店铺;3:seller;4:sku;5:category;6:shopCategory)
    ParticipateType   int64 `json:"participate_type,omitempty"`

}
