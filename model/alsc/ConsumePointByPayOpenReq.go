package alsc

// ConsumePointByPayOpenReq 
/* model for simplify = false
type ConsumePointByPayOpenReq struct {

    // 品牌id
    
    BrandId   string `json:"brand_id,omitempty"`
    

    // 变更积分
    
    ChangePoint   int64 `json:"change_point,omitempty"`
    

    // 顾客id
    
    CustomerId   string `json:"customer_id,omitempty"`
    

    // 门店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 外部门店ID
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 外部品牌id,brandId与out_brand_id不可同时为空
    
    OutBrandId   string `json:"out_brand_id,omitempty"`
    

}
*/

// ConsumePointByPayOpenReq 
type ConsumePointByPayOpenReq struct {

    // 品牌id
    BrandId   string `json:"brand_id,omitempty"`

    // 变更积分
    ChangePoint   int64 `json:"change_point,omitempty"`

    // 顾客id
    CustomerId   string `json:"customer_id,omitempty"`

    // 门店id
    ShopId   string `json:"shop_id,omitempty"`

    // 外部门店ID
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 外部品牌id,brandId与out_brand_id不可同时为空
    OutBrandId   string `json:"out_brand_id,omitempty"`

}
