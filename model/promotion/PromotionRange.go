package promotion

// PromotionRange 
/* model for simplify = false
type PromotionRange struct {

    // 参与活动的商品id。
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 活动id。
    
    ActivityId   int64 `json:"activity_id,omitempty"`
    

    // 活动名称。
    
    ActivityName   string `json:"activity_name,omitempty"`
    

}
*/

// PromotionRange 
type PromotionRange struct {

    // 参与活动的商品id。
    ItemId   int64 `json:"item_id,omitempty"`

    // 活动id。
    ActivityId   int64 `json:"activity_id,omitempty"`

    // 活动名称。
    ActivityName   string `json:"activity_name,omitempty"`

}
