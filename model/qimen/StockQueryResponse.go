package qimen

// StockQueryResponse 
type StockQueryResponse struct {
    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // 响应码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 响应信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 总数
    TotalCount   int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
    // 商品的库存信息列表
    Items   []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
