package alitripmerchant

import (
	"sync"
)

// UserAgreementVo 结构体
type UserAgreementVo struct {
	// 隐私协议
	ShowPrivacyAgreement bool `json:"show_privacy_agreement,omitempty" xml:"show_privacy_agreement,omitempty"`
	// 境外输出个人信息协议
	ShowDataExportAgreement bool `json:"show_data_export_agreement,omitempty" xml:"show_data_export_agreement,omitempty"`
}

var poolUserAgreementVo = sync.Pool{
	New: func() any {
		return new(UserAgreementVo)
	},
}

// GetUserAgreementVo() 从对象池中获取UserAgreementVo
func GetUserAgreementVo() *UserAgreementVo {
	return poolUserAgreementVo.Get().(*UserAgreementVo)
}

// ReleaseUserAgreementVo 释放UserAgreementVo
func ReleaseUserAgreementVo(v *UserAgreementVo) {
	v.ShowPrivacyAgreement = false
	v.ShowDataExportAgreement = false
	poolUserAgreementVo.Put(v)
}
