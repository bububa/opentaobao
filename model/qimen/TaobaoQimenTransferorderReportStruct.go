package qimen

// TaobaoQimenTransferorderReportStruct 
/* model for simplify = false
type TaobaoQimenTransferorderReportStruct struct {

    // 调拨单号,0,string(50),必填,
    
    TransferOrderCode   string `json:"transferOrderCode,omitempty"`
    

    // 调拨出库单号
    
    TransferOutOrderCode   string `json:"transferOutOrderCode,omitempty"`
    

    // 调拨入库单号
    
    TransferInOrderCode   string `json:"transferInOrderCode,omitempty"`
    

    // 确认出库时间
    
    ConfirmOutTime   string `json:"confirmOutTime,omitempty"`
    

    // 确认入库时间
    
    ConfirmInTime   string `json:"confirmInTime,omitempty"`
    

    // 调拨单创建时间
    
    CreateTime   string `json:"createTime,omitempty"`
    

    // 调拨出库仓编码
    
    FromWarehouseCode   string `json:"fromWarehouseCode,omitempty"`
    

    // 调拨入库仓编码
    
    ToWarehouseCode   string `json:"toWarehouseCode,omitempty"`
    

    // 111
    
    OwnerCode   string `json:"ownerCode,omitempty"`
    

    // erpOrderCode
    
    ErpOrderCode   string `json:"erpOrderCode,omitempty"`
    

    // orderStatus
    
    OrderStatus   string `json:"orderStatus,omitempty"`
    

    // 项目集
    
    Items  struct {
        Items  []Items `json:"items,omitempty"`
    } `json:"items,omitempty"`
    

    // 响应结果:success|failure,success,string(10),必填,
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应码,0,string(50),,
    
    Code   string `json:"code,omitempty"`
    

    // 响应信息,invalid appkey,string(100),,
    
    Message   string `json:"message,omitempty"`
    

}
*/

// TaobaoQimenTransferorderReportStruct 
type TaobaoQimenTransferorderReportStruct struct {

    // 调拨单号,0,string(50),必填,
    TransferOrderCode   string `json:"transferOrderCode,omitempty"`

    // 调拨出库单号
    TransferOutOrderCode   string `json:"transferOutOrderCode,omitempty"`

    // 调拨入库单号
    TransferInOrderCode   string `json:"transferInOrderCode,omitempty"`

    // 确认出库时间
    ConfirmOutTime   string `json:"confirmOutTime,omitempty"`

    // 确认入库时间
    ConfirmInTime   string `json:"confirmInTime,omitempty"`

    // 调拨单创建时间
    CreateTime   string `json:"createTime,omitempty"`

    // 调拨出库仓编码
    FromWarehouseCode   string `json:"fromWarehouseCode,omitempty"`

    // 调拨入库仓编码
    ToWarehouseCode   string `json:"toWarehouseCode,omitempty"`

    // 111
    OwnerCode   string `json:"ownerCode,omitempty"`

    // erpOrderCode
    ErpOrderCode   string `json:"erpOrderCode,omitempty"`

    // orderStatus
    OrderStatus   string `json:"orderStatus,omitempty"`

    // 项目集
    Items   []Items `json:"items,omitempty"`

    // 响应结果:success|failure,success,string(10),必填,
    Flag   string `json:"flag,omitempty"`

    // 响应码,0,string(50),,
    Code   string `json:"code,omitempty"`

    // 响应信息,invalid appkey,string(100),,
    Message   string `json:"message,omitempty"`

}
