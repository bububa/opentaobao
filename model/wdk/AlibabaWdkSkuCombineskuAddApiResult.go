package wdk

// AlibabawdkskucombineskuaddApiResult 结构体
type AlibabawdkskucombineskuaddApiResult struct {
	// 单个商品新建异常编码（异常才有值）
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 单个商品新建异常信息（异常才有值）
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 商品code值
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 单个商品是否新建成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
