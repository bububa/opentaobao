package qimen

// TaobaoQimenOrderexceptionReportRequest 
type TaobaoQimenOrderexceptionReportRequest struct {
    // 奇门仓储字段
    MessageId   string `json:"messageId,omitempty" xml:"messageId,omitempty"`
    // 奇门仓储字段
    WarehouseCode   string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
    // 奇门仓储字段
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
    // 奇门仓储字段
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
    // 奇门仓储字段
    OrderType   string `json:"orderType,omitempty" xml:"orderType,omitempty"`
    // 奇门仓储字段
    LogisticsCode   string `json:"logisticsCode,omitempty" xml:"logisticsCode,omitempty"`
    // 奇门仓储字段
    ExpressCode   string `json:"expressCode,omitempty" xml:"expressCode,omitempty"`
    // 奇门仓储字段
    MessageType   string `json:"messageType,omitempty" xml:"messageType,omitempty"`
    // 奇门仓储字段
    MessageDesc   string `json:"messageDesc,omitempty" xml:"messageDesc,omitempty"`
    // 奇门仓储字段
    CreateTime   string `json:"createTime,omitempty" xml:"createTime,omitempty"`
    // 奇门仓储字段
    OrderLines   []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
    // 奇门仓储字段
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 扩展属性
    ExtendProps   *TaobaoQimenOrderexceptionReportMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
