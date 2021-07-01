package nazca

// ChangeAuthCallBackDo 
type ChangeAuthCallBackDo struct {
    // 身份证号
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    // 支付宝账号
    AlipayAccount   string `json:"alipay_account,omitempty" xml:"alipay_account,omitempty"`
    // 邮箱
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    // 姓名
    PersonName   string `json:"person_name,omitempty" xml:"person_name,omitempty"`
    // 客户在1688的唯一标识
    PlatformUserId   string `json:"platform_user_id,omitempty" xml:"platform_user_id,omitempty"`
    // 认证状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 联系手机
    Telphone   string `json:"telphone,omitempty" xml:"telphone,omitempty"`
}
