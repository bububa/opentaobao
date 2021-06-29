package maitix

// MoaOrderContactInfo 
type MoaOrderContactInfo struct {

    // 联系人姓名-必填
    
    ContactName   string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
    

    // 联系人国家码-必填
    
    CountryCode   string `json:"country_code,omitempty" xml:"country_code,omitempty"`
    

    // 联系人email
    
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    

    // 联系人手机号-必填
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

}
