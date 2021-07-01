package alihealth2

// AlibabaAlihealthDentalStatementQueryMtopResult 结构体
type AlibabaAlihealthDentalStatementQueryMtopResult struct {
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// model
	StatementDetailVos []StatementDetailVo `json:"statement_detail_vos,omitempty" xml:"statement_detail_vos>statement_detail_vo,omitempty"`
}
