package wdk

// OfferLicenseInfo 
/* model for simplify = false
type OfferLicenseInfo struct {

    // 是否首选证件
    
    First   string `json:"first,omitempty"`
    

    // 证件开始时间
    
    GmtBegin   string `json:"gmt_begin,omitempty"`
    

    // 证件结束日期
    
    GmtValidity   string `json:"gmt_validity,omitempty"`
    

    // 证件国家
    
    LicenseCountry   string `json:"license_country,omitempty"`
    

    // 证件类型
    
    LicenseType   string `json:"license_type,omitempty"`
    

    // 证件号码
    
    LicenseValue   string `json:"license_value,omitempty"`
    

}
*/

// OfferLicenseInfo 
type OfferLicenseInfo struct {

    // 是否首选证件
    First   string `json:"first,omitempty"`

    // 证件开始时间
    GmtBegin   string `json:"gmt_begin,omitempty"`

    // 证件结束日期
    GmtValidity   string `json:"gmt_validity,omitempty"`

    // 证件国家
    LicenseCountry   string `json:"license_country,omitempty"`

    // 证件类型
    LicenseType   string `json:"license_type,omitempty"`

    // 证件号码
    LicenseValue   string `json:"license_value,omitempty"`

}
