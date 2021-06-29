package qimen

// DeliveryOrderBatchCreateRequest 
type DeliveryOrderBatchCreateRequest struct {

    // 订单信息
    
    Orders   []Order `json:"orders,omitempty" xml:"orders,omitempty"`
    

    // 扩展属性
    
    ExtendProps   *TaobaoQimenDeliveryorderBatchcreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
    

}
