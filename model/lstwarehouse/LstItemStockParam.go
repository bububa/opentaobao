package lstwarehouse

// LstItemStockParam 
type LstItemStockParam struct {
    // 库存参数列表
    StockList   []Stocklist `json:"stock_list,omitempty" xml:"stock_list>stocklist,omitempty"`
}
