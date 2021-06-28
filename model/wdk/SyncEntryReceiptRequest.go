package wdk

// SyncEntryReceiptRequest 
/* model for simplify = false
type SyncEntryReceiptRequest struct {

    // 地址信息
    
    AddressInfo  struct {
        AddressInfo  []AddressInfo `json:"address_info,omitempty"`
    } `json:"address_info,omitempty"`
    

    // 联系人信息
    
    ContactInfo  struct {
        ContactInfo  []ContactInfo `json:"contact_info,omitempty"`
    } `json:"contact_info,omitempty"`
    

    // 教育经历
    
    EducationExpInfo  struct {
        EducationExpInfo  []EducationExpInfo `json:"education_exp_info,omitempty"`
    } `json:"education_exp_info,omitempty"`
    

    // 员工主要信息
    
    EmployeeBasic  *struct {
        EmployeeBasic  *EmployeeBasic `json:"employee_basic,omitempty"`
    } `json:"employee_basic,omitempty"`
    

    // 家庭信息
    
    FamilyInfo  struct {
        FamilyInfo  []FamilyInfo `json:"family_info,omitempty"`
    } `json:"family_info,omitempty"`
    

    // 工作履历
    
    JobExpInfo  struct {
        JobExpInfo  []JobExpInfo `json:"job_exp_info,omitempty"`
    } `json:"job_exp_info,omitempty"`
    

    // 证件信息
    
    OfferLicenseInfo  struct {
        OfferLicenseInfo  []OfferLicenseInfo `json:"offer_license_info,omitempty"`
    } `json:"offer_license_info,omitempty"`
    

    // 关联方申报信息
    
    RelatedPartyInfo  struct {
        RelatedPartyInfo  []RelatedPartyInfo `json:"related_party_info,omitempty"`
    } `json:"related_party_info,omitempty"`
    

    // 语言信息
    
    LanguageInfo  struct {
        LanguageInfo  []LanguageInfo `json:"language_info,omitempty"`
    } `json:"language_info,omitempty"`
    

}
*/

// SyncEntryReceiptRequest 
type SyncEntryReceiptRequest struct {

    // 地址信息
    AddressInfo   []AddressInfo `json:"address_info,omitempty"`

    // 联系人信息
    ContactInfo   []ContactInfo `json:"contact_info,omitempty"`

    // 教育经历
    EducationExpInfo   []EducationExpInfo `json:"education_exp_info,omitempty"`

    // 员工主要信息
    EmployeeBasic   *EmployeeBasic `json:"employee_basic,omitempty"`

    // 家庭信息
    FamilyInfo   []FamilyInfo `json:"family_info,omitempty"`

    // 工作履历
    JobExpInfo   []JobExpInfo `json:"job_exp_info,omitempty"`

    // 证件信息
    OfferLicenseInfo   []OfferLicenseInfo `json:"offer_license_info,omitempty"`

    // 关联方申报信息
    RelatedPartyInfo   []RelatedPartyInfo `json:"related_party_info,omitempty"`

    // 语言信息
    LanguageInfo   []LanguageInfo `json:"language_info,omitempty"`

}
