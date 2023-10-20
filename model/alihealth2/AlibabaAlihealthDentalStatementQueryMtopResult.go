package alihealth2

import (
	"sync"
)

// AlibabaAlihealthDentalStatementQueryMtopResult 结构体
type AlibabaAlihealthDentalStatementQueryMtopResult struct {
	// model
	StatementDetailVos []StatementDetailVo `json:"statement_detail_vos,omitempty" xml:"statement_detail_vos>statement_detail_vo,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDentalStatementQueryMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalStatementQueryMtopResult)
	},
}

// GetAlibabaAlihealthDentalStatementQueryMtopResult() 从对象池中获取AlibabaAlihealthDentalStatementQueryMtopResult
func GetAlibabaAlihealthDentalStatementQueryMtopResult() *AlibabaAlihealthDentalStatementQueryMtopResult {
	return poolAlibabaAlihealthDentalStatementQueryMtopResult.Get().(*AlibabaAlihealthDentalStatementQueryMtopResult)
}

// ReleaseAlibabaAlihealthDentalStatementQueryMtopResult 释放AlibabaAlihealthDentalStatementQueryMtopResult
func ReleaseAlibabaAlihealthDentalStatementQueryMtopResult(v *AlibabaAlihealthDentalStatementQueryMtopResult) {
	v.StatementDetailVos = v.StatementDetailVos[:0]
	v.Success = false
	poolAlibabaAlihealthDentalStatementQueryMtopResult.Put(v)
}
