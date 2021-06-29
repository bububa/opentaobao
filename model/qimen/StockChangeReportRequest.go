package qimen

// StockChangeReportRequest 
type StockChangeReportRequest struct {

    // item
    
    Items   []Item `json:"items,omitempty" xml:"items,omitempty"`
    

    // 扩展属性
    
    ExtendProps   *TaobaoQimenStockchangeReportMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
    

}
