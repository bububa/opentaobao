package wdk

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
