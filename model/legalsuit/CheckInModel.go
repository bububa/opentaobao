package legalsuit

import (
	"sync"
)

// CheckInModel 结构体
type CheckInModel struct {
	// 内部承办人
	UndertakerPeoples []string `json:"undertaker_peoples,omitempty" xml:"undertaker_peoples>string,omitempty"`
	// 管辖异议截止时间 根据收案时间来定
	JurisdictionObjectionDeadlineTime string `json:"jurisdiction_objection_deadline_time,omitempty" xml:"jurisdiction_objection_deadline_time,omitempty"`
	// 举证截止时间 未来要自动计算
	ProofDeadlineTime string `json:"proof_deadline_time,omitempty" xml:"proof_deadline_time,omitempty"`
	// 法务收案时间
	LegalCheckinTime string `json:"legal_checkin_time,omitempty" xml:"legal_checkin_time,omitempty"`
}

var poolCheckInModel = sync.Pool{
	New: func() any {
		return new(CheckInModel)
	},
}

// GetCheckInModel() 从对象池中获取CheckInModel
func GetCheckInModel() *CheckInModel {
	return poolCheckInModel.Get().(*CheckInModel)
}

// ReleaseCheckInModel 释放CheckInModel
func ReleaseCheckInModel(v *CheckInModel) {
	v.UndertakerPeoples = v.UndertakerPeoples[:0]
	v.JurisdictionObjectionDeadlineTime = ""
	v.ProofDeadlineTime = ""
	v.LegalCheckinTime = ""
	poolCheckInModel.Put(v)
}
