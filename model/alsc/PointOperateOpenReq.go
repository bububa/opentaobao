package alsc

// PointOperateOpenReq 
/* model for simplify = false
type PointOperateOpenReq struct {

    // saas品牌id
    
    BrandId   string `json:"brand_id,omitempty"`
    

    // 操作积分
    
    ChangePoint   int64 `json:"change_point,omitempty"`
    

    // 手机号
    
    Mobile   string `json:"mobile,omitempty"`
    

    // 1-增加(charge)  2-冻结(freeze)  3-核销(verify)  4-扣减(decrease)
    
    OperateType   int64 `json:"operate_type,omitempty"`
    

    // 操作人id
    
    OperatorId   string `json:"operator_id,omitempty"`
    

    // 外部id
    
    OuterId   string `json:"outer_id,omitempty"`
    

    // 外部id类型，wechat：微信openId  alipay：支付宝
    
    OuterType   string `json:"outer_type,omitempty"`
    

    // 操作原因
    
    Reason   string `json:"reason,omitempty"`
    

    // 请求幂等id
    
    RequestId   string `json:"request_id,omitempty"`
    

    // saas门店id
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 外部订单号
    
    OuterOrderId   string `json:"outer_order_id,omitempty"`
    

    // 顾客id
    
    CustomerId   string `json:"customer_id,omitempty"`
    

}
*/

// PointOperateOpenReq 
type PointOperateOpenReq struct {

    // saas品牌id
    BrandId   string `json:"brand_id,omitempty"`

    // 操作积分
    ChangePoint   int64 `json:"change_point,omitempty"`

    // 手机号
    Mobile   string `json:"mobile,omitempty"`

    // 1-增加(charge)  2-冻结(freeze)  3-核销(verify)  4-扣减(decrease)
    OperateType   int64 `json:"operate_type,omitempty"`

    // 操作人id
    OperatorId   string `json:"operator_id,omitempty"`

    // 外部id
    OuterId   string `json:"outer_id,omitempty"`

    // 外部id类型，wechat：微信openId  alipay：支付宝
    OuterType   string `json:"outer_type,omitempty"`

    // 操作原因
    Reason   string `json:"reason,omitempty"`

    // 请求幂等id
    RequestId   string `json:"request_id,omitempty"`

    // saas门店id
    ShopId   string `json:"shop_id,omitempty"`

    // 外部订单号
    OuterOrderId   string `json:"outer_order_id,omitempty"`

    // 顾客id
    CustomerId   string `json:"customer_id,omitempty"`

}
