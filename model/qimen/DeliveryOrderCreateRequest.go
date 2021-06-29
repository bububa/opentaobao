package qimen

// DeliveryOrderCreateRequest 
type DeliveryOrderCreateRequest struct {

    // 发货单信息
    
    DeliveryOrder   *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
    

    // 订单列表
    
    OrderLines   []OrderLine `json:"orderLines,omitempty" xml:"orderLines,omitempty"`
    

    // 扩展属性
    
    ExtendProps   *TaobaoQimenDeliveryorderCreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
    

}
