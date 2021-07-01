package category

// WholesaleCategoryOpenResult 结构体
type WholesaleCategoryOpenResult struct {
	// result_memo 返回结果描述 例如 “success”
	ResultMemo string `json:"result_memo,omitempty" xml:"result_memo,omitempty"`
	// result_code 返回码 例如200表示成功
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 具体类目信息
	Result *WholesaleCategory `json:"result,omitempty" xml:"result,omitempty"`
}
