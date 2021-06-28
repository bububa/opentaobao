package alsc

// CalculateDeductedMoneyOpenReq 
type CalculateDeductedMoneyOpenReq struct {

    // 品牌id
    
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    

    // 抵现积分数量
    
    ConsumePoint   int64 `json:"consume_point,omitempty" xml:"consume_point,omitempty"`
    

    // 顾客id
    
    CustomerId   string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
    

    // 门店id
    
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    

    // 外部门店ID,shop_id和out_shop_id不可同时为空
    
    OutShopId   string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
    

    // 外部品牌id,brandId与out_brand_id不可同时为空
    
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    

}
