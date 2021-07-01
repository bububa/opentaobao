package product

// WholesaleSearchOpenResult 结构体
type WholesaleSearchOpenResult struct {
	// result_memo 返回结果描述，例如success表示成功
	ResultMemo string `json:"result_memo,omitempty" xml:"result_memo,omitempty"`
	// result_code 返回结果码，例如200
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 在线批发产品搜索结果
	Result *GoodsSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}
