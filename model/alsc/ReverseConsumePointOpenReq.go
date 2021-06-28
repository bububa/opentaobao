package alsc

// ReverseConsumePointOpenReq 
/* model for simplify = false
type ReverseConsumePointOpenReq struct {

    // 业务场景
    
    BizType   string `json:"biz_type,omitempty"`
    

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
    

    // 变更理由
    
    Reason   string `json:"reason,omitempty"`
    

    // 请求幂等UUID
    
    RequestId   string `json:"request_id,omitempty"`
    

    // 反向操作的业务ID,如退单号
    
    ReverseOrderId   string `json:"reverse_order_id,omitempty"`
    

    // 门店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 外部门店ID,shop_id和out_shop_id不可同时为空
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 外部品牌id,brandId与out_brand_id不可同时为空
    
    OutBrandId   string `json:"out_brand_id,omitempty"`
    

    // 外部渠道，CS是辰森，客如云是KRY
    
    BizChannel   string `json:"biz_channel,omitempty"`
    

}
*/

// ReverseConsumePointOpenReq 
type ReverseConsumePointOpenReq struct {

    // 业务场景
    BizType   string `json:"biz_type,omitempty"`

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

    // 变更理由
    Reason   string `json:"reason,omitempty"`

    // 请求幂等UUID
    RequestId   string `json:"request_id,omitempty"`

    // 反向操作的业务ID,如退单号
    ReverseOrderId   string `json:"reverse_order_id,omitempty"`

    // 门店id
    ShopId   string `json:"shop_id,omitempty"`

    // 外部门店ID,shop_id和out_shop_id不可同时为空
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 外部品牌id,brandId与out_brand_id不可同时为空
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 外部渠道，CS是辰森，客如云是KRY
    BizChannel   string `json:"biz_channel,omitempty"`

}
