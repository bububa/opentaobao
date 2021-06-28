package qimen

// EntryOrderConfirmRequest 
/* model for simplify = false
type EntryOrderConfirmRequest struct {

    // 入库单信息
    
    EntryOrder  *struct {
        EntryOrder  *EntryOrder `json:"entry_order,omitempty"`
    } `json:"entryOrder,omitempty"`
    

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

// EntryOrderConfirmRequest 
type EntryOrderConfirmRequest struct {

    // 入库单信息
    EntryOrder   *EntryOrder `json:"entryOrder,omitempty"`

    // 订单信息
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
