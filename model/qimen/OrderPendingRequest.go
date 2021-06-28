package qimen

// OrderPendingRequest 
/* model for simplify = false
type OrderPendingRequest struct {

    // 操作类型(pending=挂起;restore=恢复)
    
    ActionType   string `json:"actionType,omitempty"`
    

    // 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    
    WarehouseCode   string `json:"warehouseCode,omitempty"`
    

    // 货主编码
    
    OwnerCode   string `json:"ownerCode,omitempty"`
    

    // 单据编码
    
    OrderCode   string `json:"orderCode,omitempty"`
    

    // 仓储系统单据编码
    
    OrderId   string `json:"orderId,omitempty"`
    

    // 单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK=换货入库;CNJG=仓内加工单)
    
    OrderType   string `json:"orderType,omitempty"`
    

    // 挂起/恢复原因
    
    Reason   string `json:"reason,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

}
*/

// OrderPendingRequest 
type OrderPendingRequest struct {

    // 操作类型(pending=挂起;restore=恢复)
    ActionType   string `json:"actionType,omitempty"`

    // 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 货主编码
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 单据编码
    OrderCode   string `json:"orderCode,omitempty"`

    // 仓储系统单据编码
    OrderId   string `json:"orderId,omitempty"`

    // 单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK=换货入库;CNJG=仓内加工单)
    OrderType   string `json:"orderType,omitempty"`

    // 挂起/恢复原因
    Reason   string `json:"reason,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
