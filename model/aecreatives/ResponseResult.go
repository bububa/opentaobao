package aecreatives

// ResponseResult 结构体
type ResponseResult struct {
	// 返回状态描述
	RespMsg string `json:"resp_msg,omitempty" xml:"resp_msg,omitempty"`
	// 返回状态码
	RespCode int64 `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// 返回记录结果列表
	Result *AliexpressAffiliateCategoryGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
