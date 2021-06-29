package qimen

// DeliveryOrderCreateResponse 
type DeliveryOrderCreateResponse struct {
    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // 响应码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 响应信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 订单创建时间(YYYY-MM-DD HH:MM:SS)
    CreateTime   string `json:"createTime,omitempty" xml:"createTime,omitempty"`
    // 出库单仓储系统编码
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
    // 仓库编码(统仓统配使用)
    WarehouseCode   string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
    // 物流公司编码(统仓统配使用)
    LogisticsCode   string `json:"logisticsCode,omitempty" xml:"logisticsCode,omitempty"`
    // 发货单信息
    DeliveryOrders   []DeliveryOrder `json:"deliveryOrders,omitempty" xml:"deliveryOrders>delivery_order,omitempty"`
}
