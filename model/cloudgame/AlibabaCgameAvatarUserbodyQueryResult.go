package cloudgame

// AlibabacgameavataruserbodyqueryResult 结构体
type AlibabacgameavataruserbodyqueryResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应数据
	Data *AlibabacgameavataruserbodyqueryData `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
