package qimen

// EntryOrderCreateRequest 
/* model for simplify = false
type EntryOrderCreateRequest struct {

    // 入库单信息
    
    EntryOrder  *struct {
        EntryOrder  *EntryOrder `json:"entry_order,omitempty"`
    } `json:"entryOrder,omitempty"`
    

    // 入库单详情
    
    OrderLines  struct {
        OrderLine  []OrderLine `json:"order_line,omitempty"`
    } `json:"orderLines,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

}
*/

// EntryOrderCreateRequest 
type EntryOrderCreateRequest struct {

    // 入库单信息
    EntryOrder   *EntryOrder `json:"entryOrder,omitempty"`

    // 入库单详情
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
