package qimen

// DeliveryOrderCreateResponse 
type DeliveryOrderCreateResponse struct {

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 订单创建时间(YYYY-MM-DD HH:MM:SS)
    CreateTime   string `json:"createTime,omitempty"`

    // 出库单仓储系统编码
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty"`

    // 仓库编码(统仓统配使用)
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 物流公司编码(统仓统配使用)
    LogisticsCode   string `json:"logisticsCode,omitempty"`

    // 发货单信息
    DeliveryOrders   []DeliveryOrder `json:"deliveryOrders,omitempty"`

}
