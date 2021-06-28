package alsc

// BatchActiveCardOpenReq 
/* model for simplify = false
type BatchActiveCardOpenReq struct {

    // 品牌id
    
    BrandId   string `json:"brand_id,omitempty"`
    

    // 操作员id
    
    OperatorId   string `json:"operator_id,omitempty"`
    

    // 外部品牌id,brandId与out_brand_id不可同时为空
    
    OutBrandId   string `json:"out_brand_id,omitempty"`
    

    // 外部门店ID,shop_id和out_shop_id不可同时为空
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 实体卡列表
    
    PhysicalCardIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"physical_card_ids,omitempty"`
    

    // 请求id
    
    RequestId   string `json:"request_id,omitempty"`
    

    // 门店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

}
*/

// BatchActiveCardOpenReq 
type BatchActiveCardOpenReq struct {

    // 品牌id
    BrandId   string `json:"brand_id,omitempty"`

    // 操作员id
    OperatorId   string `json:"operator_id,omitempty"`

    // 外部品牌id,brandId与out_brand_id不可同时为空
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 外部门店ID,shop_id和out_shop_id不可同时为空
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 实体卡列表
    PhysicalCardIds   []string `json:"physical_card_ids,omitempty"`

    // 请求id
    RequestId   string `json:"request_id,omitempty"`

    // 门店id
    ShopId   string `json:"shop_id,omitempty"`

}
