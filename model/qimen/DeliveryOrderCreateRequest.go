package qimen

// DeliveryOrderCreateRequest 
type DeliveryOrderCreateRequest struct {

    // 发货单信息
    DeliveryOrder   *DeliveryOrder `json:"deliveryOrder,omitempty"`

    // 订单列表
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
