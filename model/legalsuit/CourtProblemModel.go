package legalsuit

import (
	"sync"
)

// CourtProblemModel 结构体
type CourtProblemModel struct {
	// 当事人回复
	PartyReply string `json:"party_reply,omitempty" xml:"party_reply,omitempty"`
	// 问题描述
	ProblemDescription string `json:"problem_description,omitempty" xml:"problem_description,omitempty"`
	// 被询问人
	Askeder string `json:"askeder,omitempty" xml:"askeder,omitempty"`
	// 问题是否有价值
	IsCount string `json:"is_count,omitempty" xml:"is_count,omitempty"`
}

var poolCourtProblemModel = sync.Pool{
	New: func() any {
		return new(CourtProblemModel)
	},
}

// GetCourtProblemModel() 从对象池中获取CourtProblemModel
func GetCourtProblemModel() *CourtProblemModel {
	return poolCourtProblemModel.Get().(*CourtProblemModel)
}

// ReleaseCourtProblemModel 释放CourtProblemModel
func ReleaseCourtProblemModel(v *CourtProblemModel) {
	v.PartyReply = ""
	v.ProblemDescription = ""
	v.Askeder = ""
	v.IsCount = ""
	poolCourtProblemModel.Put(v)
}
