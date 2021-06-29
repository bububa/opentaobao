package alsc

// UpdatePayPasswdReq 
type UpdatePayPasswdReq struct {
    // 品牌id  、 外部品牌id    2选1
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    // 修改人名称
    OperatorName   string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
    // 操作人id
    OperatorId   string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    // md5代表密码已经MD5加密，raw代表密码是原生没有加密
    PasswdType   string `json:"passwd_type,omitempty" xml:"passwd_type,omitempty"`
    // 密码
    PayPasswd   string `json:"pay_passwd,omitempty" xml:"pay_passwd,omitempty"`
    // 顾客id
    CustomerId   string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
    // 外部品牌id
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    // 请求id
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
