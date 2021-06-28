package qimen

// EntryOrderQueryResponse 
/* model for simplify = false
type EntryOrderQueryResponse struct {

    // 响应结果:success|failure
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应码
    
    Code   string `json:"code,omitempty"`
    

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

    // orderLines总条数
    
    TotalLines   int64 `json:"totalLines,omitempty"`
    

    // 入库单信息
    
    EntryOrder  *struct {
        EntryOrder  *EntryOrder `json:"entry_order,omitempty"`
    } `json:"entryOrder,omitempty"`
    

    // 入库单单据信息
    
    OrderLines  struct {
        OrderLine  []OrderLine `json:"order_line,omitempty"`
    } `json:"orderLines,omitempty"`
    

}
*/

// EntryOrderQueryResponse 
type EntryOrderQueryResponse struct {

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

    // orderLines总条数
    TotalLines   int64 `json:"totalLines,omitempty"`

    // 入库单信息
    EntryOrder   *EntryOrder `json:"entryOrder,omitempty"`

    // 入库单单据信息
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

}
