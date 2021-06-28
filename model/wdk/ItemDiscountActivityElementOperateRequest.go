package wdk

// ItemDiscountActivityElementOperateRequest 
/* model for simplify = false
type ItemDiscountActivityElementOperateRequest struct {

    // 商品元素信息
    
    SkuElements  struct {
        SkuActivityElementDto  []SkuActivityElementDto `json:"sku_activity_element_dto,omitempty"`
    } `json:"sku_elements,omitempty"`
    

    // 操作活动的ID
    
    ActId   int64 `json:"act_id,omitempty"`
    

    // 操作人ID（仅支持数字类型）
    
    CreatorId   string `json:"creator_id,omitempty"`
    

    // 操作人Name
    
    CreatorName   string `json:"creator_name,omitempty"`
    

}
*/

// ItemDiscountActivityElementOperateRequest 
type ItemDiscountActivityElementOperateRequest struct {

    // 商品元素信息
    SkuElements   []SkuActivityElementDto `json:"sku_elements,omitempty"`

    // 操作活动的ID
    ActId   int64 `json:"act_id,omitempty"`

    // 操作人ID（仅支持数字类型）
    CreatorId   string `json:"creator_id,omitempty"`

    // 操作人Name
    CreatorName   string `json:"creator_name,omitempty"`

}
