package alihealthoutflow

import (
	"sync"
)

// SyncVisitInfoRequest 结构体
type SyncVisitInfoRequest struct {
	// 业务创建时间(非空)
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 业务类型(非空)
	BusinessType string `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 业务id(非空)
	BusinessId string `json:"business_id,omitempty" xml:"business_id,omitempty"`
	// 授权id(非空)
	AuthorizationId string `json:"authorization_id,omitempty" xml:"authorization_id,omitempty"`
	// 医院id(非空)
	HospitalId string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
	// 医生id(非空)
	DoctorId string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
	// 医生名(非空)
	DoctorName string `json:"doctor_name,omitempty" xml:"doctor_name,omitempty"`
	// 科室id(非空)
	DepartId string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// 科室名(非空)
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 患者id(非空)
	PatientId string `json:"patient_id,omitempty" xml:"patient_id,omitempty"`
	// 患者名(非空)
	PatientName string `json:"patient_name,omitempty" xml:"patient_name,omitempty"`
	// 状态(非空)
	VisitStatus string `json:"visit_status,omitempty" xml:"visit_status,omitempty"`
	// 首次回复时间(可空)
	FirstResponseTime string `json:"first_response_time,omitempty" xml:"first_response_time,omitempty"`
}

var poolSyncVisitInfoRequest = sync.Pool{
	New: func() any {
		return new(SyncVisitInfoRequest)
	},
}

// GetSyncVisitInfoRequest() 从对象池中获取SyncVisitInfoRequest
func GetSyncVisitInfoRequest() *SyncVisitInfoRequest {
	return poolSyncVisitInfoRequest.Get().(*SyncVisitInfoRequest)
}

// ReleaseSyncVisitInfoRequest 释放SyncVisitInfoRequest
func ReleaseSyncVisitInfoRequest(v *SyncVisitInfoRequest) {
	v.CreateTime = ""
	v.BusinessType = ""
	v.BusinessId = ""
	v.AuthorizationId = ""
	v.HospitalId = ""
	v.DoctorId = ""
	v.DoctorName = ""
	v.DepartId = ""
	v.DepartName = ""
	v.PatientId = ""
	v.PatientName = ""
	v.VisitStatus = ""
	v.FirstResponseTime = ""
	poolSyncVisitInfoRequest.Put(v)
}
