package alsc

// BindPhysicalCardOpenReq 
type BindPhysicalCardOpenReq struct {

    // 外部品牌ID,brand_id与out_brand_id不可同时为空
    
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    

    // 物理卡号
    
    PhysicalCardId   string `json:"physical_card_id,omitempty" xml:"physical_card_id,omitempty"`
    

    // 请求ID，用于幂等处理
    
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    

    // 品牌ID,,brand_id与out_brand_id不可同时为空
    
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    

    // 外部门店ID，shop_id和out_shop_id不可同时为空
    
    OutShopId   string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
    

    // 店铺ID，shop_id和out_shop_id不可同时为空
    
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    

    // 操作人id
    
    OperatorId   string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    

    // 用户id
    
    CustomerId   string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
    

    // 卡号
    
    CardId   string `json:"card_id,omitempty" xml:"card_id,omitempty"`
    

}
