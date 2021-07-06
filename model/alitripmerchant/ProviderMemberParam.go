package alitripmerchant

// ProviderMemberParam 结构体
type ProviderMemberParam struct {
	// 用户英文姓
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// 称呼
	Civility string `json:"civility,omitempty" xml:"civility,omitempty"`
	// 签名
	Signature string `json:"signature,omitempty" xml:"signature,omitempty"`
	// 手机号前缀（加密）
	PhonePre string `json:"phone_pre,omitempty" xml:"phone_pre,omitempty"`
	// 手机号（加密）
	PhoneNum string `json:"phone_num,omitempty" xml:"phone_num,omitempty"`
	// 语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 请求路径
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 英文名
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// 请求应用的id
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 邮箱（加密）
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 追踪信息
	PromoteInfo string `json:"promote_info,omitempty" xml:"promote_info,omitempty"`
	// 是否订阅信息
	Subscription bool `json:"subscription,omitempty" xml:"subscription,omitempty"`
	// 是否接受协议
	AcceptedTandC bool `json:"accepted_tand_c,omitempty" xml:"accepted_tand_c,omitempty"`
}
