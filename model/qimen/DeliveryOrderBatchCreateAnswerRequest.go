package qimen

// DeliveryOrderBatchCreateAnswerRequest 
/* model for simplify = false
type DeliveryOrderBatchCreateAnswerRequest struct {

    // 发货单列表
    
    Orders  struct {
        Order  []Order `json:"order,omitempty"`
    } `json:"orders,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

}
*/

// DeliveryOrderBatchCreateAnswerRequest 
type DeliveryOrderBatchCreateAnswerRequest struct {

    // 发货单列表
    Orders   []Order `json:"orders,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
