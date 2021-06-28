package wdk

// BuyGiftActivitySkuOperateRequest 
type BuyGiftActivitySkuOperateRequest struct {

    // 商品元素信息
    
    SkuElements   []SkuActivityElementDto `json:"sku_elements,omitempty" xml:"sku_elements,omitempty"`
    

    // 操作活动的ID
    
    ActId   int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
    

    // 操作人ID（数字类型）
    
    CreatorId   string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
    

    // 操作人Name
    
    CreatorName   string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
    

}
