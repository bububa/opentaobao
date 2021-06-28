package wdk

// SyncEntryReceiptRequest 
type SyncEntryReceiptRequest struct {

    // 地址信息
    
    AddressInfo   []AddressInfo `json:"address_info,omitempty" xml:"address_info,omitempty"`
    

    // 联系人信息
    
    ContactInfo   []ContactInfo `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
    

    // 教育经历
    
    EducationExpInfo   []EducationExpInfo `json:"education_exp_info,omitempty" xml:"education_exp_info,omitempty"`
    

    // 员工主要信息
    
    EmployeeBasic   *EmployeeBasic `json:"employee_basic,omitempty" xml:"employee_basic,omitempty"`
    

    // 家庭信息
    
    FamilyInfo   []FamilyInfo `json:"family_info,omitempty" xml:"family_info,omitempty"`
    

    // 工作履历
    
    JobExpInfo   []JobExpInfo `json:"job_exp_info,omitempty" xml:"job_exp_info,omitempty"`
    

    // 证件信息
    
    OfferLicenseInfo   []OfferLicenseInfo `json:"offer_license_info,omitempty" xml:"offer_license_info,omitempty"`
    

    // 关联方申报信息
    
    RelatedPartyInfo   []RelatedPartyInfo `json:"related_party_info,omitempty" xml:"related_party_info,omitempty"`
    

    // 语言信息
    
    LanguageInfo   []LanguageInfo `json:"language_info,omitempty" xml:"language_info,omitempty"`
    

}
