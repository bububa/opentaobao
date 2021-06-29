package fenxiao

// RefundDetail 
type RefundDetail struct {
    // 子单id
    SubOrderId   int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
    // 是否退货
    IsReturnGoods   bool `json:"is_return_goods,omitempty" xml:"is_return_goods,omitempty"`
    // 退款创建时间
    RefundCreateTime   string `json:"refund_create_time,omitempty" xml:"refund_create_time,omitempty"`
    // 退款状态<br/>1：买家已经申请退款，等待卖家同意<br/>2：卖家已经同意退款，等待买家退货<br/>3：买家已经退货，等待卖家确认收货<br/>4：退款关闭<br/>5：退款成功<br/>6：卖家拒绝退款<br/>12：同意退款，待打款<br/>9：没有申请退款<br/>10：卖家拒绝确认收货
    RefundStatus   int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
    // 退款的金额
    RefundFee   string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
    // 支付给供应商的金额
    PaySupFee   string `json:"pay_sup_fee,omitempty" xml:"pay_sup_fee,omitempty"`
    // 退款原因
    RefundReason   string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
    // 退款说明
    RefundDesc   string `json:"refund_desc,omitempty" xml:"refund_desc,omitempty"`
    // 下游买家的退款信息
    BuyerRefund   *BuyerRefund `json:"buyer_refund,omitempty" xml:"buyer_refund,omitempty"`
    // 供应商nick
    SupplierNick   string `json:"supplier_nick,omitempty" xml:"supplier_nick,omitempty"`
    // 分销商nick
    DistributorNick   string `json:"distributor_nick,omitempty" xml:"distributor_nick,omitempty"`
    // 退款修改时间。格式:yyyy-MM-dd HH:mm:ss
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    // 主采购单id
    PurchaseOrderId   int64 `json:"purchase_order_id,omitempty" xml:"purchase_order_id,omitempty"`
    // 退款流程类型：<br/>4：发货前退款；<br/>1：发货后退款不退货；<br/>2：发货后退款退货
    RefundFlowType   int64 `json:"refund_flow_type,omitempty" xml:"refund_flow_type,omitempty"`
    // 超时时间
    Timeout   string `json:"timeout,omitempty" xml:"timeout,omitempty"`
    // 超时类型：<br/>1：供应商同意退款/同意退货超时；<br/>2：供应商确认收货超时
    ToType   int64 `json:"to_type,omitempty" xml:"to_type,omitempty"`
}
