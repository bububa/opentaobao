package qimen

// OrderCallbackRequestDo 
type OrderCallbackRequestDo struct {
    // 奇门仓储字段,C123,string(50),,
    WarehouseCode   string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
    // 奇门仓储字段,C123,string(50),,
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    // 奇门仓储字段,C123,string(50),,
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
    // 奇门仓储字段,C123,string(50),,
    OrderId   string `json:"orderId,omitempty" xml:"orderId,omitempty"`
    // 运单号
    ExpressCode   string `json:"expressCode,omitempty" xml:"expressCode,omitempty"`
}
