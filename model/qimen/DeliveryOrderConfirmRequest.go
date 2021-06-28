package qimen

// DeliveryOrderConfirmRequest 
/* model for simplify = false
type DeliveryOrderConfirmRequest struct {

    // 发货单信息
    
    DeliveryOrder  *struct {
        DeliveryOrder  *DeliveryOrder `json:"delivery_order,omitempty"`
    } `json:"deliveryOrder,omitempty"`
    

    // 包裹信息
    
    Packages  struct {
        Package  []Package `json:"package,omitempty"`
    } `json:"packages,omitempty"`
    

    // 单据列表
    
    OrderLines  struct {
        OrderLine  []OrderLine `json:"order_line,omitempty"`
    } `json:"orderLines,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

}
*/

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
