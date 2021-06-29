package alicom

// SupplierOrderResultModel 
type SupplierOrderResultModel struct {
    // 业务类型：7-合约机分销、
    BizType   string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    // 结果码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 结果描述
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    // 交易订单号
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    // 供应商外部订单号
    OutOrderNo   string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
    // 订购结果状态
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
