package alihealth

// PrescriptionDoctorAuthRequest 
/* model for simplify = false
type PrescriptionDoctorAuthRequest struct {

    // 医生id(非空)
    
    DoctorId   string `json:"doctor_id,omitempty"`
    

    // 支付宝userId(非空)
    
    AlipayUserId   string `json:"alipay_user_id,omitempty"`
    

    // 医院id(非空)
    
    HospitalId   string `json:"hospital_id,omitempty"`
    

}
*/

// PrescriptionDoctorAuthRequest 
type PrescriptionDoctorAuthRequest struct {

    // 医生id(非空)
    DoctorId   string `json:"doctor_id,omitempty"`

    // 支付宝userId(非空)
    AlipayUserId   string `json:"alipay_user_id,omitempty"`

    // 医院id(非空)
    HospitalId   string `json:"hospital_id,omitempty"`

}
