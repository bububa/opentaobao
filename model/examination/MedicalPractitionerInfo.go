package examination

// MedicalPractitionerInfo 
type MedicalPractitionerInfo struct {

    // 性别(MALE-男;FEMALE-女;)
    
    Gender   string `json:"gender,omitempty" xml:"gender,omitempty"`
    

    // 从业者手机号
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 外部系统从业者id
    
    OuterPractitionerId   string `json:"outer_practitioner_id,omitempty" xml:"outer_practitioner_id,omitempty"`
    

    // 从业者姓名
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 从业者资质证书图片链接
    
    CertificationsUrl   string `json:"certifications_url,omitempty" xml:"certifications_url,omitempty"`
    

    // 执业机构名称
    
    PracticeAgencyName   string `json:"practice_agency_name,omitempty" xml:"practice_agency_name,omitempty"`
    

    // 执业机构所在省份
    
    PracticeAgencyProvince   string `json:"practice_agency_province,omitempty" xml:"practice_agency_province,omitempty"`
    

    // 执业机构所在省份Code
    
    PracticeAgencyProvinceCode   string `json:"practice_agency_province_code,omitempty" xml:"practice_agency_province_code,omitempty"`
    

}
