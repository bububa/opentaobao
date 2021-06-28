package wdk

// ItemPoolActivityElementOperateRequest 
/* model for simplify = false
type ItemPoolActivityElementOperateRequest struct {

    // 商品元素列表
    
    SkuElements  struct {
        ItemPoolSkuActivityElementDto  []ItemPoolSkuActivityElementDto `json:"item_pool_sku_activity_element_dto,omitempty"`
    } `json:"sku_elements,omitempty"`
    

    // 同城零售活动id
    
    ActId   int64 `json:"act_id,omitempty"`
    

    // 操作人id
    
    CreatorId   string `json:"creator_id,omitempty"`
    

    // 操作人名称
    
    CreatorName   string `json:"creator_name,omitempty"`
    

    // 商品所属分组序号，默认单个分组则填1
    
    GroupNumber   int64 `json:"group_number,omitempty"`
    

}
*/

// ItemPoolActivityElementOperateRequest 
type ItemPoolActivityElementOperateRequest struct {

    // 商品元素列表
    SkuElements   []ItemPoolSkuActivityElementDto `json:"sku_elements,omitempty"`

    // 同城零售活动id
    ActId   int64 `json:"act_id,omitempty"`

    // 操作人id
    CreatorId   string `json:"creator_id,omitempty"`

    // 操作人名称
    CreatorName   string `json:"creator_name,omitempty"`

    // 商品所属分组序号，默认单个分组则填1
    GroupNumber   int64 `json:"group_number,omitempty"`

}
