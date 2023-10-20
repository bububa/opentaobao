package alihouse

// TmallalihousetradecouponordercodeexchangeResult 结构体
type TmallalihousetradecouponordercodeexchangeResult struct {
	// 结果说明
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
