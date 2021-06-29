package qimen

// TransferOrderDetail 
type TransferOrderDetail struct {
    // 调拨单号,0,string(50),,
    TransferOrderCode   string `json:"transferOrderCode,omitempty" xml:"transferOrderCode,omitempty"`
    // 外部ERP订单号,HZ1234,string(50),,
    ErpOrderCode   string `json:"erpOrderCode,omitempty" xml:"erpOrderCode,omitempty"`
    // 订单状态,0,string(50),,
    OrderStatus   string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty"`
    // 调拨出库单号
    TransferOutOrderCode   string `json:"transferOutOrderCode,omitempty" xml:"transferOutOrderCode,omitempty"`
    // 调拨入库单号
    TransferInOrderCode   string `json:"transferInOrderCode,omitempty" xml:"transferInOrderCode,omitempty"`
    // 创建时间
    CreateTime   string `json:"createTime,omitempty" xml:"createTime,omitempty"`
    // 确认出库时间
    ConfirmOutTime   string `json:"confirmOutTime,omitempty" xml:"confirmOutTime,omitempty"`
    // 确认入库时间
    ConfirmInTime   string `json:"confirmInTime,omitempty" xml:"confirmInTime,omitempty"`
    // 调拨出库仓编码
    FromWarehouseCode   string `json:"fromWarehouseCode,omitempty" xml:"fromWarehouseCode,omitempty"`
    // 调拨入库仓编码
    ToWarehouseCode   string `json:"toWarehouseCode,omitempty" xml:"toWarehouseCode,omitempty"`
    // 1111
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    // 调拨单货品明细记录集
    Items   []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
