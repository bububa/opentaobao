package qimen

// DeliveryOrderBatchCreateAnswerRequest 
type DeliveryOrderBatchCreateAnswerRequest struct {

    // 发货单列表
    
    Orders   []Order `json:"orders,omitempty" xml:"orders,omitempty"`
    

    // 扩展属性
    
    ExtendProps   *TaobaoQimenDeliveryorderBatchcreateAnswerMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
    

}
