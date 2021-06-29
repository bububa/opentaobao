package wdk

// FinanceOrderDetailRequest 
type FinanceOrderDetailRequest struct {
    // 当前页码
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    // 每页条数
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 销售渠道
    SaleChannel   string `json:"sale_channel,omitempty" xml:"sale_channel,omitempty"`
    // 销售来源
    SaleSource   string `json:"sale_source,omitempty" xml:"sale_source,omitempty"`
    // 交易类型
    TradeType   string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
    // 门店编码list
    ShopCodes   []Null `json:"shop_codes,omitempty" xml:"shop_codes>null,omitempty"`
    // 结束时间
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 开始时间
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
}
