package alihealthoutflow

// SyncPatientInfoRequest 
type SyncPatientInfoRequest struct {
    // 患者id(非空)
    PatientId   string `json:"patient_id,omitempty" xml:"patient_id,omitempty"`
    // 患者姓名(非空)
    PatientName   string `json:"patient_name,omitempty" xml:"patient_name,omitempty"`
    // 身份证号(非空)
    IdCard   string `json:"id_card,omitempty" xml:"id_card,omitempty"`
    // 患者手机号(非空)
    MobilePhone   string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
    // 所属医院id(非空)
    HospitalId   string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
}
