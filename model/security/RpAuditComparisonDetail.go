package security

import (
	"sync"
)

// RpAuditComparisonDetail 结构体
type RpAuditComparisonDetail struct {
	// score
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// resultType
	ResultType *RpAuditType `json:"result_type,omitempty" xml:"result_type,omitempty"`
	// value
	Value *RpAuditValue `json:"value,omitempty" xml:"value,omitempty"`
}

var poolRpAuditComparisonDetail = sync.Pool{
	New: func() any {
		return new(RpAuditComparisonDetail)
	},
}

// GetRpAuditComparisonDetail() 从对象池中获取RpAuditComparisonDetail
func GetRpAuditComparisonDetail() *RpAuditComparisonDetail {
	return poolRpAuditComparisonDetail.Get().(*RpAuditComparisonDetail)
}

// ReleaseRpAuditComparisonDetail 释放RpAuditComparisonDetail
func ReleaseRpAuditComparisonDetail(v *RpAuditComparisonDetail) {
	v.Score = ""
	v.ResultType = nil
	v.Value = nil
	poolRpAuditComparisonDetail.Put(v)
}
