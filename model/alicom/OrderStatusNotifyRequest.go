package alicom

// OrderStatusNotifyRequest 
type OrderStatusNotifyRequest struct {

    // 订单状态
    Orderstatus   bool `json:"orderstatus,omitempty"`

    // 错误描述
    ErrDesc   string `json:"err_desc,omitempty"`

    // 卖家id
    SellerId   string `json:"seller_id,omitempty"`

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 业务类型
    BizType   string `json:"biz_type,omitempty"`

    // 天猫订单号
    BizOrderId   string `json:"biz_order_id,omitempty"`

    // 外部流水号
    OutSerialNumber   string `json:"out_serial_number,omitempty"`

}
