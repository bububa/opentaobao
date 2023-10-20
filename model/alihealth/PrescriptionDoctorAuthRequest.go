package alihealth

import (
	"sync"
)

// PrescriptionDoctorAuthRequest 结构体
type PrescriptionDoctorAuthRequest struct {
	// 医生id(非空)
	DoctorId string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
	// 支付宝userId(非空)
	AlipayUserId string `json:"alipay_user_id,omitempty" xml:"alipay_user_id,omitempty"`
	// 医院id(非空)
	HospitalId string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
}

var poolPrescriptionDoctorAuthRequest = sync.Pool{
	New: func() any {
		return new(PrescriptionDoctorAuthRequest)
	},
}

// GetPrescriptionDoctorAuthRequest() 从对象池中获取PrescriptionDoctorAuthRequest
func GetPrescriptionDoctorAuthRequest() *PrescriptionDoctorAuthRequest {
	return poolPrescriptionDoctorAuthRequest.Get().(*PrescriptionDoctorAuthRequest)
}

// ReleasePrescriptionDoctorAuthRequest 释放PrescriptionDoctorAuthRequest
func ReleasePrescriptionDoctorAuthRequest(v *PrescriptionDoctorAuthRequest) {
	v.DoctorId = ""
	v.AlipayUserId = ""
	v.HospitalId = ""
	poolPrescriptionDoctorAuthRequest.Put(v)
}
