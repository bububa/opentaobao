package alitripmerchant

// QueryMemberParam 
type QueryMemberParam struct {
    // 签名
    Signature   string `json:"signature,omitempty" xml:"signature,omitempty"`
    // 手机号前缀（加密）
    PhonePre   string `json:"phone_pre,omitempty" xml:"phone_pre,omitempty"`
    // 手机号（加密）
    PhoneNum   string `json:"phone_num,omitempty" xml:"phone_num,omitempty"`
    // 应用id
    AppId   string `json:"app_id,omitempty" xml:"app_id,omitempty"`
}
