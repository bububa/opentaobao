package alsc

// OrderBackflowOpenReq 
type OrderBackflowOpenReq struct {

    // 待支付动作-WAIT_PAY，  支付动作-PAID,  订单成功动作-SUCCESS，  订单关闭动作-CLOSE，  退款动作-REFUND
    Action   string `json:"action,omitempty"`

    // 品牌id
    BrandId   string `json:"brand_id,omitempty"`

    // 扩展
    ExtInfo   string `json:"ext_info,omitempty"`

    // 发票信息
    InvoiceInfo   *InvoiceInfo `json:"invoice_info,omitempty"`

    // 物流信息
    LogisticalInfo   *LogisticalInfo `json:"logistical_info,omitempty"`

    // 操作人
    OperatorId   string `json:"operator_id,omitempty"`

    // 订单信息
    OrderInfo   *OrderInfo `json:"order_info,omitempty"`

    // 支付信息
    PaymentInfoList   []PaymentInfo `json:"payment_info_list,omitempty"`

    // 退款信息
    RefundInfoList   []RefundInfo `json:"refund_info_list,omitempty"`

    // 请求id
    RequestId   string `json:"request_id,omitempty"`

    // 门店id
    ShopId   string `json:"shop_id,omitempty"`

    // 操作人姓名
    OperatorName   string `json:"operator_name,omitempty"`

    // 服务人员类别
    ServicePersonCategory   string `json:"service_person_category,omitempty"`

    // 服务人员名称
    ServicePersonName   string `json:"service_person_name,omitempty"`

    // 服务人员ID
    ServicePersonId   string `json:"service_person_id,omitempty"`

    // 订单类型-餐饮-RESTAURANT  生服-LIFE_SERVICE
    OrderChannel   string `json:"order_channel,omitempty"`

}
