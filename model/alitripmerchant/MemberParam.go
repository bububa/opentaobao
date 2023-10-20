package alitripmerchant

import (
	"sync"
)

// MemberParam 结构体
type MemberParam struct {
	// 用户英文姓
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// 用户英文名
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// 称呼
	Civility string `json:"civility,omitempty" xml:"civility,omitempty"`
	// 用户手机号前缀
	PhonePre string `json:"phone_pre,omitempty" xml:"phone_pre,omitempty"`
	// 用户语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 用户手机号
	PhoneNum string `json:"phone_num,omitempty" xml:"phone_num,omitempty"`
	// 用户邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 城市
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 用户注册类型（1.新用户注册  2.老用户注册）
	UserRegisterType string `json:"user_register_type,omitempty" xml:"user_register_type,omitempty"`
	// 是否接受协议
	AcceptedTandC bool `json:"accepted_tand_c,omitempty" xml:"accepted_tand_c,omitempty"`
	// 是否接受消息推送
	Subscription bool `json:"subscription,omitempty" xml:"subscription,omitempty"`
	// 用户是否勾选短信收取优惠邮件
	OptinAll bool `json:"optin_all,omitempty" xml:"optin_all,omitempty"`
	// 是否统一向境外提供个人信息
	DataExportAgreement bool `json:"data_export_agreement,omitempty" xml:"data_export_agreement,omitempty"`
}

var poolMemberParam = sync.Pool{
	New: func() any {
		return new(MemberParam)
	},
}

// GetMemberParam() 从对象池中获取MemberParam
func GetMemberParam() *MemberParam {
	return poolMemberParam.Get().(*MemberParam)
}

// ReleaseMemberParam 释放MemberParam
func ReleaseMemberParam(v *MemberParam) {
	v.LastName = ""
	v.FirstName = ""
	v.Civility = ""
	v.PhonePre = ""
	v.Language = ""
	v.PhoneNum = ""
	v.Email = ""
	v.Country = ""
	v.UserRegisterType = ""
	v.AcceptedTandC = false
	v.Subscription = false
	v.OptinAll = false
	v.DataExportAgreement = false
	poolMemberParam.Put(v)
}
