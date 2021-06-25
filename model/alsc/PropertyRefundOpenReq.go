package alsc

// PropertyRefundOpenReq 
type PropertyRefundOpenReq struct {

    // saas品牌id
    BrandId   string `json:"brand_id,omitempty"`

    // 手机号
    Mobile   string `json:"mobile,omitempty"`

    // 退款的订单号
    NewOuterOrderId   string `json:"new_outer_order_id,omitempty"`

    // 操作人
    OperatorId   string `json:"operator_id,omitempty"`

    // 外部id
    OuterId   string `json:"outer_id,omitempty"`

    // 原订单号，就是核销的订单号
    OuterOrderId   string `json:"outer_order_id,omitempty"`

    // 外部id类型，wechat微信，alipay支付宝
    OuterType   string `json:"outer_type,omitempty"`

    // 请求幂等id
    RequestId   string `json:"request_id,omitempty"`

    // saas门店id
    ShopId   string `json:"shop_id,omitempty"`

    // 券实例id集合
    VoucherIdList   []String `json:"voucher_id_list,omitempty"`

    // 顾客id
    CustomerId   string `json:"customer_id,omitempty"`

}
