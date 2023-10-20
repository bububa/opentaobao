package legalsuit

import (
	"sync"
)

// CourtEvidenceModel 结构体
type CourtEvidenceModel struct {
	// 质证意见
	Opinion string `json:"opinion,omitempty" xml:"opinion,omitempty"`
	// 是否认可
	IsAgreed string `json:"is_agreed,omitempty" xml:"is_agreed,omitempty"`
	// 质证人
	Evidence string `json:"evidence,omitempty" xml:"evidence,omitempty"`
	// 证据目的
	EvidenceAim string `json:"evidence_aim,omitempty" xml:"evidence_aim,omitempty"`
	// 证据名称
	EvidenceName string `json:"evidence_name,omitempty" xml:"evidence_name,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 质证证据类型(原告证据还是被告证据)
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolCourtEvidenceModel = sync.Pool{
	New: func() any {
		return new(CourtEvidenceModel)
	},
}

// GetCourtEvidenceModel() 从对象池中获取CourtEvidenceModel
func GetCourtEvidenceModel() *CourtEvidenceModel {
	return poolCourtEvidenceModel.Get().(*CourtEvidenceModel)
}

// ReleaseCourtEvidenceModel 释放CourtEvidenceModel
func ReleaseCourtEvidenceModel(v *CourtEvidenceModel) {
	v.Opinion = ""
	v.IsAgreed = ""
	v.Evidence = ""
	v.EvidenceAim = ""
	v.EvidenceName = ""
	v.Name = ""
	v.Type = ""
	poolCourtEvidenceModel.Put(v)
}
