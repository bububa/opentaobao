package qimen

// TaobaoQimenTransferorderReportStruct 
type TaobaoQimenTransferorderReportStruct struct {
    // 调拨单号,0,string(50),必填,
    TransferOrderCode   string `json:"transferOrderCode,omitempty" xml:"transferOrderCode,omitempty"`
    // 调拨出库单号
    TransferOutOrderCode   string `json:"transferOutOrderCode,omitempty" xml:"transferOutOrderCode,omitempty"`
    // 调拨入库单号
    TransferInOrderCode   string `json:"transferInOrderCode,omitempty" xml:"transferInOrderCode,omitempty"`
    // 确认出库时间
    ConfirmOutTime   string `json:"confirmOutTime,omitempty" xml:"confirmOutTime,omitempty"`
    // 确认入库时间
    ConfirmInTime   string `json:"confirmInTime,omitempty" xml:"confirmInTime,omitempty"`
    // 调拨单创建时间
    CreateTime   string `json:"createTime,omitempty" xml:"createTime,omitempty"`
    // 调拨出库仓编码
    FromWarehouseCode   string `json:"fromWarehouseCode,omitempty" xml:"fromWarehouseCode,omitempty"`
    // 调拨入库仓编码
    ToWarehouseCode   string `json:"toWarehouseCode,omitempty" xml:"toWarehouseCode,omitempty"`
    // 111
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    // erpOrderCode
    ErpOrderCode   string `json:"erpOrderCode,omitempty" xml:"erpOrderCode,omitempty"`
    // orderStatus
    OrderStatus   string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty"`
    // 项目集
    Items   []Items `json:"items,omitempty" xml:"items>items,omitempty"`
    // 响应结果:success|failure,success,string(10),必填,
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // 响应码,0,string(50),,
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 响应信息,invalid appkey,string(100),,
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
