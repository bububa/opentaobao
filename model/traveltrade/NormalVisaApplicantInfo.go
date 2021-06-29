package traveltrade

// NormalVisaApplicantInfo 
type NormalVisaApplicantInfo struct {
    // 可选，申请人ID。更新申请人基本信息时必填，新增申请人信息时不用填
    ApplyId   string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 姓，新增申请人时必填
    Surname   string `json:"surname,omitempty" xml:"surname,omitempty"`
    // 名，新增申请人时必填
    GivenName   string `json:"given_name,omitempty" xml:"given_name,omitempty"`
    // 手机号，新增申请人时必填
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    // 证件号（护照、入台证等），新增申请人时必填
    CertNo   string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
    // 办理人身份类型(8-在职人员,9-自由职业,10-在校学生,11-退休人员,12-学龄年儿童,13-所有申请者,14-单个成年人,15-随行直系亲属,16-在读学生)
    UserType   int64 `json:"user_type,omitempty" xml:"user_type,omitempty"`
    // 办理人中文全名
    ApplyNameCn   string `json:"apply_name_cn,omitempty" xml:"apply_name_cn,omitempty"`
}
