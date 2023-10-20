package security

import (
	"sync"
)

// RpAuditComparisonDetailBos 结构体
type RpAuditComparisonDetailBos struct {
	// 比对分值
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 比对结论
	Value *RpAuditValueBo `json:"value,omitempty" xml:"value,omitempty"`
	// 比对项
	ResultType *RpAuditTypeBo `json:"result_type,omitempty" xml:"result_type,omitempty"`
}

var poolRpAuditComparisonDetailBos = sync.Pool{
	New: func() any {
		return new(RpAuditComparisonDetailBos)
	},
}

// GetRpAuditComparisonDetailBos() 从对象池中获取RpAuditComparisonDetailBos
func GetRpAuditComparisonDetailBos() *RpAuditComparisonDetailBos {
	return poolRpAuditComparisonDetailBos.Get().(*RpAuditComparisonDetailBos)
}

// ReleaseRpAuditComparisonDetailBos 释放RpAuditComparisonDetailBos
func ReleaseRpAuditComparisonDetailBos(v *RpAuditComparisonDetailBos) {
	v.Score = ""
	v.Value = nil
	v.ResultType = nil
	poolRpAuditComparisonDetailBos.Put(v)
}
