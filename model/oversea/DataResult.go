package oversea

// DataResult 
type DataResult struct {
    // 查到的税率信息
    ExchangeRate   string `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
    // 查询结果是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误代码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 错误信息
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
}
