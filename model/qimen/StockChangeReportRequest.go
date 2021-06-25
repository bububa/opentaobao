package qimen

// StockChangeReportRequest 
type StockChangeReportRequest struct {

    // item
    Items   []Item `json:"items,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
