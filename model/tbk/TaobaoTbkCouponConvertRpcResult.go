package tbk

// TaobaotbkcouponconvertRpcResult 结构体
type TaobaotbkcouponconvertRpcResult struct {
	// 见错误描述
	Resultmsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// data
	Results *MaterialDto `json:"results,omitempty" xml:"results,omitempty"`
	// 见错误码描述
	Resultcode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
