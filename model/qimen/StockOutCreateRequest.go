package qimen

// StockOutCreateRequest 
/* model for simplify = false
type StockOutCreateRequest struct {

    // 出库单信息
    
    DeliveryOrder  *struct {
        DeliveryOrder  *DeliveryOrder `json:"delivery_order,omitempty"`
    } `json:"deliveryOrder,omitempty"`
    

    // 单据信息
    
    OrderLines  struct {
        OrderLine  []OrderLine `json:"order_line,omitempty"`
    } `json:"orderLines,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

}
*/

// StockOutCreateRequest 
type StockOutCreateRequest struct {

    // 出库单信息
    DeliveryOrder   *DeliveryOrder `json:"deliveryOrder,omitempty"`

    // 单据信息
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
