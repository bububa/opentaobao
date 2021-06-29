package alidoc

// RxPrescriptionTopDTO 
type RxPrescriptionTopDTO struct {
    // 药品列表
    DrugList   []RxDrugTopDTO `json:"drug_list,omitempty" xml:"drug_list>rx_drug_top_dto,omitempty"`
    // 患者问诊信息
    PatientDiagnostic   *RxPatientDiagnosticTopDTO `json:"patient_diagnostic,omitempty" xml:"patient_diagnostic,omitempty"`
    // 患者信息
    Patient   *RxPatientTopDTO `json:"patient,omitempty" xml:"patient,omitempty"`
    // 处方图片url
    PicUrl   string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
    // 处方创建时间
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 处方id
    RxId   string `json:"rx_id,omitempty" xml:"rx_id,omitempty"`
    // 医生信息
    Doctor   *RxDoctorTopDTO `json:"doctor,omitempty" xml:"doctor,omitempty"`
}
