package qimen

// StockChangeReportRequest 
/* model for simplify = false
type StockChangeReportRequest struct {

    // item
    
    Items  struct {
        Item  []Item `json:"item,omitempty"`
    } `json:"items,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

}
*/

// StockChangeReportRequest 
type StockChangeReportRequest struct {

    // item
    Items   []Item `json:"items,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
