package alihouse

// AlibabaAlihouseNewhomeShopconfigAstorePreviewResult 结构体
type AlibabaAlihouseNewhomeShopconfigAstorePreviewResult struct {
	// 返回的结果
	Data []AstoreRespDto `json:"data,omitempty" xml:"data>astore_resp_dto,omitempty"`
	// 处理完成后的消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误/成功码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
