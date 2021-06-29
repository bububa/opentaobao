package dutyfree

// AlibabaDutyfreeStockQueryResultDO 
type AlibabaDutyfreeStockQueryResultDO struct {
    // 错误码
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // 调用是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 具体库存信息
    Object   *StockResultDTO `json:"object,omitempty" xml:"object,omitempty"`
}
