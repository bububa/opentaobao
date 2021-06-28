package qimen

// DeliveryOrderBatchConfirmRequest 
/* model for simplify = false
type DeliveryOrderBatchConfirmRequest struct {

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

// DeliveryOrderBatchConfirmRequest 
type DeliveryOrderBatchConfirmRequest struct {

    // 发货单列表
    Orders   []Order `json:"orders,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
