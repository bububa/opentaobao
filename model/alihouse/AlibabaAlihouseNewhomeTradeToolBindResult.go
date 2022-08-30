package alihouse

// AlibabaAlihouseNewhomeTradeToolBindResult 结构体
type AlibabaAlihouseNewhomeTradeToolBindResult struct {
	// data
	Data []TradeToolBindResultDto `json:"data,omitempty" xml:"data>trade_tool_bind_result_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
