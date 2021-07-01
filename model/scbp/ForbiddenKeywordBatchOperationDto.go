package scbp

// ForbiddenKeywordBatchOperationDto 结构体
type ForbiddenKeywordBatchOperationDto struct {
	// 请求参数
	ForbiddenKeywordOperationList []ForbiddenKeywordOperation `json:"forbidden_keyword_operation_list,omitempty" xml:"forbidden_keyword_operation_list>forbidden_keyword_operation,omitempty"`
}
