package alitripmerchant

// MemberBaseInfoDto 结构体
type MemberBaseInfoDto struct {
	// 用户英文姓
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// 用户英文名
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// 称呼
	Civility string `json:"civility,omitempty" xml:"civility,omitempty"`
	// 语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
}
