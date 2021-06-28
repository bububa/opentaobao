package alsc

// ConsumePointOpenReq 
/* model for simplify = false
type ConsumePointOpenReq struct {

    // 品牌id
    
    BrandId   string `json:"brand_id,omitempty"`
    

    // 变更积分数
    
    ChangePoint   int64 `json:"change_point,omitempty"`
    

    // 顾客id
    
    CustomerId   string `json:"customer_id,omitempty"`
    

    // 操作人id
    
    OperatorId   string `json:"operator_id,omitempty"`
    

    // 操作人名
    
    OperatorName   string `json:"operator_name,omitempty"`
    

    // 关联交易号/订单号
    
    OutBizId   string `json:"out_biz_id,omitempty"`
    

    // 变更原因
    
    Reason   string `json:"reason,omitempty"`
    

    // 请求id，用来做幂等
    
    RequestId   string `json:"request_id,omitempty"`
    

    // 门店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 外部门店ID,shop_id和out_shop_id不可同时为空
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 外部品牌id,brandId与out_brand_id不可同时为空
    
    OutBrandId   string `json:"out_brand_id,omitempty"`
    

    // 辰森是CS,客如云是KRY
    
    BizChannel   string `json:"biz_channel,omitempty"`
    

}
*/

// ConsumePointOpenReq 
type ConsumePointOpenReq struct {

    // 品牌id
    BrandId   string `json:"brand_id,omitempty"`

    // 变更积分数
    ChangePoint   int64 `json:"change_point,omitempty"`

    // 顾客id
    CustomerId   string `json:"customer_id,omitempty"`

    // 操作人id
    OperatorId   string `json:"operator_id,omitempty"`

    // 操作人名
    OperatorName   string `json:"operator_name,omitempty"`

    // 关联交易号/订单号
    OutBizId   string `json:"out_biz_id,omitempty"`

    // 变更原因
    Reason   string `json:"reason,omitempty"`

    // 请求id，用来做幂等
    RequestId   string `json:"request_id,omitempty"`

    // 门店id
    ShopId   string `json:"shop_id,omitempty"`

    // 外部门店ID,shop_id和out_shop_id不可同时为空
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 外部品牌id,brandId与out_brand_id不可同时为空
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 辰森是CS,客如云是KRY
    BizChannel   string `json:"biz_channel,omitempty"`

}
