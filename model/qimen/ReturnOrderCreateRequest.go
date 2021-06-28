package qimen

// ReturnOrderCreateRequest 
/* model for simplify = false
type ReturnOrderCreateRequest struct {

    // 退货单信息
    
    ReturnOrder  *struct {
        ReturnOrder  *ReturnOrder `json:"return_order,omitempty"`
    } `json:"returnOrder,omitempty"`
    

    // 订单信息
    
    OrderLines  struct {
        OrderLine  []OrderLine `json:"order_line,omitempty"`
    } `json:"orderLines,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

}
*/

// ReturnOrderCreateRequest 
type ReturnOrderCreateRequest struct {

    // 退货单信息
    ReturnOrder   *ReturnOrder `json:"returnOrder,omitempty"`

    // 订单信息
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
