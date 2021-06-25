package qimen

// OrderCallbackRequestDO 
type OrderCallbackRequestDO struct {

    // 奇门仓储字段,C123,string(50),,
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 奇门仓储字段,C123,string(50),,
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 奇门仓储字段,C123,string(50),,
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty"`

    // 奇门仓储字段,C123,string(50),,
    OrderId   string `json:"orderId,omitempty"`

    // 运单号
    ExpressCode   string `json:"expressCode,omitempty"`

}
