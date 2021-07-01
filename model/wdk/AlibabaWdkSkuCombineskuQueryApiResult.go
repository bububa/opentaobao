package wdk

// AlibabaWdkSkuCombineskuQueryApiResult 结构体
type AlibabaWdkSkuCombineskuQueryApiResult struct {
	// 单个商品是否查询成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 单个商品查询异常编码（异常才有值）
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 单个商品查询异常信息（异常才有值）
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 商品信息
	Model *SkuDo `json:"model,omitempty" xml:"model,omitempty"`
}
