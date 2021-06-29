package qimen

// DeliveryOrderBatchConfirmRequest 
type DeliveryOrderBatchConfirmRequest struct {

    // 发货单列表
    
    Orders   []Order `json:"orders,omitempty" xml:"orders,omitempty"`
    

    // 扩展属性
    
    ExtendProps   *TaobaoQimenDeliveryorderBatchconfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
    

}
