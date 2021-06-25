package qimen

// DeliveryOrderConfirmRequest 
type DeliveryOrderConfirmRequest struct {

    // 发货单信息
    DeliveryOrder   *DeliveryOrder `json:"deliveryOrder,omitempty"`

    // 包裹信息
    Packages   []Package `json:"packages,omitempty"`

    // 单据列表
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
