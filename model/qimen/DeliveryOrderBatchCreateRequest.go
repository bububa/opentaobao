package qimen

// DeliveryOrderBatchCreateRequest 
/* model for simplify = false
type DeliveryOrderBatchCreateRequest struct {

    // 订单信息
    
    Orders  struct {
        Order  []Order `json:"order,omitempty"`
    } `json:"orders,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

}
*/

// DeliveryOrderBatchCreateRequest 
type DeliveryOrderBatchCreateRequest struct {

    // 订单信息
    Orders   []Order `json:"orders,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
