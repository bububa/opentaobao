package qimen

// TransferOrderDetail 
/* model for simplify = false
type TransferOrderDetail struct {

    // 调拨单号,0,string(50),,
    
    TransferOrderCode   string `json:"transferOrderCode,omitempty"`
    

    // 外部ERP订单号,HZ1234,string(50),,
    
    ErpOrderCode   string `json:"erpOrderCode,omitempty"`
    

    // 订单状态,0,string(50),,
    
    OrderStatus   string `json:"orderStatus,omitempty"`
    

    // 调拨出库单号
    
    TransferOutOrderCode   string `json:"transferOutOrderCode,omitempty"`
    

    // 调拨入库单号
    
    TransferInOrderCode   string `json:"transferInOrderCode,omitempty"`
    

    // 创建时间
    
    CreateTime   string `json:"createTime,omitempty"`
    

    // 确认出库时间
    
    ConfirmOutTime   string `json:"confirmOutTime,omitempty"`
    

    // 确认入库时间
    
    ConfirmInTime   string `json:"confirmInTime,omitempty"`
    

    // 调拨出库仓编码
    
    FromWarehouseCode   string `json:"fromWarehouseCode,omitempty"`
    

    // 调拨入库仓编码
    
    ToWarehouseCode   string `json:"toWarehouseCode,omitempty"`
    

    // 1111
    
    OwnerCode   string `json:"ownerCode,omitempty"`
    

    // 调拨单货品明细记录集
    
    Items  struct {
        Item  []Item `json:"item,omitempty"`
    } `json:"items,omitempty"`
    

}
*/

// TransferOrderDetail 
type TransferOrderDetail struct {

    // 调拨单号,0,string(50),,
    TransferOrderCode   string `json:"transferOrderCode,omitempty"`

    // 外部ERP订单号,HZ1234,string(50),,
    ErpOrderCode   string `json:"erpOrderCode,omitempty"`

    // 订单状态,0,string(50),,
    OrderStatus   string `json:"orderStatus,omitempty"`

    // 调拨出库单号
    TransferOutOrderCode   string `json:"transferOutOrderCode,omitempty"`

    // 调拨入库单号
    TransferInOrderCode   string `json:"transferInOrderCode,omitempty"`

    // 创建时间
    CreateTime   string `json:"createTime,omitempty"`

    // 确认出库时间
    ConfirmOutTime   string `json:"confirmOutTime,omitempty"`

    // 确认入库时间
    ConfirmInTime   string `json:"confirmInTime,omitempty"`

    // 调拨出库仓编码
    FromWarehouseCode   string `json:"fromWarehouseCode,omitempty"`

    // 调拨入库仓编码
    ToWarehouseCode   string `json:"toWarehouseCode,omitempty"`

    // 1111
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 调拨单货品明细记录集
    Items   []Item `json:"items,omitempty"`

}
