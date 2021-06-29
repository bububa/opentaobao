package tmallchannel

// TopChannelSubDeliverOrderDto 
type TopChannelSubDeliverOrderDto struct {
    // 子发货单单号
    SubDeliverOrderNo   int64 `json:"sub_deliver_order_no,omitempty" xml:"sub_deliver_order_no,omitempty"`
    // 主采购单号
    MainPurchaseOrderNo   int64 `json:"main_purchase_order_no,omitempty" xml:"main_purchase_order_no,omitempty"`
    // 子采购单单号
    SubPurchaseOrderNo   int64 `json:"sub_purchase_order_no,omitempty" xml:"sub_purchase_order_no,omitempty"`
    // 发货数量
    DeliverCount   int64 `json:"deliver_count,omitempty" xml:"deliver_count,omitempty"`
    // 实收数量
    ReceiveQuantity   int64 `json:"receive_quantity,omitempty" xml:"receive_quantity,omitempty"`
    // 差异的数量
    DiffQuantity   int64 `json:"diff_quantity,omitempty" xml:"diff_quantity,omitempty"`
    // 主发货单单号
    MainDeliverOrderNo   int64 `json:"main_deliver_order_no,omitempty" xml:"main_deliver_order_no,omitempty"`
}
