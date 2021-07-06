package promotion

// BenefitSendResultExt 结构体
type BenefitSendResultExt struct {
	// 异常码列表
	ErrorCodeList []string `json:"error_code_list,omitempty" xml:"error_code_list>string,omitempty"`
	// 活动详情id
	IndexId int64 `json:"index_id,omitempty" xml:"index_id,omitempty"`
	// 失败数
	FailureCount int64 `json:"failure_count,omitempty" xml:"failure_count,omitempty"`
	// 成功数
	SuccessCount int64 `json:"success_count,omitempty" xml:"success_count,omitempty"`
}
