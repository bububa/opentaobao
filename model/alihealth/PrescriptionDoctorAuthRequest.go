package alihealth

// PrescriptionDoctorAuthRequest 结构体
type PrescriptionDoctorAuthRequest struct {
	// 医生id(非空)
	DoctorId string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
	// 支付宝userId(非空)
	AlipayUserId string `json:"alipay_user_id,omitempty" xml:"alipay_user_id,omitempty"`
	// 医院id(非空)
	HospitalId string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
}
