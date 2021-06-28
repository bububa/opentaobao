package wdk

// OfferLicenseInfo 
type OfferLicenseInfo struct {

    // 是否首选证件
    
    First   string `json:"first,omitempty" xml:"first,omitempty"`
    

    // 证件开始时间
    
    GmtBegin   string `json:"gmt_begin,omitempty" xml:"gmt_begin,omitempty"`
    

    // 证件结束日期
    
    GmtValidity   string `json:"gmt_validity,omitempty" xml:"gmt_validity,omitempty"`
    

    // 证件国家
    
    LicenseCountry   string `json:"license_country,omitempty" xml:"license_country,omitempty"`
    

    // 证件类型
    
    LicenseType   string `json:"license_type,omitempty" xml:"license_type,omitempty"`
    

    // 证件号码
    
    LicenseValue   string `json:"license_value,omitempty" xml:"license_value,omitempty"`
    

}
