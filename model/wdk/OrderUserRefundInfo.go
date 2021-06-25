package wdk

// OrderUserRefundInfo 
type OrderUserRefundInfo struct {

    // 退款原因
    RefundReason   string `json:"refund_reason,omitempty"`

    // 用户备注
    Memo   string `json:"memo,omitempty"`

    // 外部退款批次Id（确保唯一，可取UUID）
    OutRefundBatchId   string `json:"out_refund_batch_id,omitempty"`

    // 盒马主单号
    BizOrderId   string `json:"biz_order_id,omitempty"`

    // 渠道店Id
    ShopId   string `json:"shop_id,omitempty"`

    // 经营店Id
    StoreId   string `json:"store_id,omitempty"`

    // 退款子单
    SubRefundOrders   []SubRefundOrder `json:"sub_refund_orders,omitempty"`

}
