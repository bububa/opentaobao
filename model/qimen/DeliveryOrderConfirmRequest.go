package qimen

// DeliveryOrderConfirmRequest 
type DeliveryOrderConfirmRequest struct {

    // 发货单信息
    
    DeliveryOrder   *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
    

    // 包裹信息
    
    Packages   []Package `json:"packages,omitempty" xml:"packages,omitempty"`
    

    // 单据列表
    
    OrderLines   []OrderLine `json:"orderLines,omitempty" xml:"orderLines,omitempty"`
    

    // 扩展属性
    
    ExtendProps   *TaobaoQimenDeliveryorderConfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
    

}
