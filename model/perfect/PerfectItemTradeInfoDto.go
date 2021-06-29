package perfect

// PerfectItemTradeInfoDTO 
type PerfectItemTradeInfoDTO struct {
    // 是否提供发票
    HasInvoice   bool `json:"has_invoice,omitempty" xml:"has_invoice,omitempty"`
    // 是否支持退换货承诺
    SellPromise   bool `json:"sell_promise,omitempty" xml:"sell_promise,omitempty"`
    // 是否拍下减库存
    SubStockAtBuy   bool `json:"sub_stock_at_buy,omitempty" xml:"sub_stock_at_buy,omitempty"`
    // 是否支持7天无理由
    Support7Day   bool `json:"support7_day,omitempty" xml:"support7_day,omitempty"`
    // 是否保修
    Warrant   bool `json:"warrant,omitempty" xml:"warrant,omitempty"`
}
