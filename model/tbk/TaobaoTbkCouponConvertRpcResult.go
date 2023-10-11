package tbk

// TaobaoTbkCouponConvertRpcResult 结构体
type TaobaoTbkCouponConvertRpcResult struct {
	// 见错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// data
	Results *MaterialDto `json:"results,omitempty" xml:"results,omitempty"`
	// 见错误码描述
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
