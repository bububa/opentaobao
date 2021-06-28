package scbp

// TopP4pModifyQuickCampaignProductDto 
/* model for simplify = false
type TopP4pModifyQuickCampaignProductDto struct {

    // 操作类型，0=商品暂停，1=商品开启，2=新增商品，3=删除商品
    
    Action   int64 `json:"action,omitempty"`
    

    // 计划ID
    
    CampaignId   int64 `json:"campaign_id,omitempty"`
    

    // 商品ID
    
    ProductIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"product_id_list,omitempty"`
    

}
*/

// TopP4pModifyQuickCampaignProductDto 
type TopP4pModifyQuickCampaignProductDto struct {

    // 操作类型，0=商品暂停，1=商品开启，2=新增商品，3=删除商品
    Action   int64 `json:"action,omitempty"`

    // 计划ID
    CampaignId   int64 `json:"campaign_id,omitempty"`

    // 商品ID
    ProductIdList   []int64 `json:"product_id_list,omitempty"`

}
