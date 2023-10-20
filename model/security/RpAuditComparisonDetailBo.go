package security

import (
	"sync"
)

// RpAuditComparisonDetailBo 结构体
type RpAuditComparisonDetailBo struct {
	// 分数
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 对比结果
	Value *RpAuditValueBo `json:"value,omitempty" xml:"value,omitempty"`
	// 类型
	ResultType *RpAuditTypeBo `json:"result_type,omitempty" xml:"result_type,omitempty"`
}

var poolRpAuditComparisonDetailBo = sync.Pool{
	New: func() any {
		return new(RpAuditComparisonDetailBo)
	},
}

// GetRpAuditComparisonDetailBo() 从对象池中获取RpAuditComparisonDetailBo
func GetRpAuditComparisonDetailBo() *RpAuditComparisonDetailBo {
	return poolRpAuditComparisonDetailBo.Get().(*RpAuditComparisonDetailBo)
}

// ReleaseRpAuditComparisonDetailBo 释放RpAuditComparisonDetailBo
func ReleaseRpAuditComparisonDetailBo(v *RpAuditComparisonDetailBo) {
	v.Score = ""
	v.Value = nil
	v.ResultType = nil
	poolRpAuditComparisonDetailBo.Put(v)
}
