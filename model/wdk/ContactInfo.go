package wdk

// ContactInfo 
type ContactInfo struct {

    // 联系人姓名
    ContactName   string `json:"contact_name,omitempty"`

    // 联系人类型
    ContactType   string `json:"contact_type,omitempty"`

    // 联系人手机号
    Mobile   string `json:"mobile,omitempty"`

}
