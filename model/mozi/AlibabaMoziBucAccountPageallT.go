package mozi

import (
	"sync"
)

// AlibabaMoziBucAccountPageallT 结构体
type AlibabaMoziBucAccountPageallT struct {
	// 账号名
	Account string `json:"account,omitempty" xml:"account,omitempty"`
	// 账号平台
	Namespace string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// 三方账号ID
	ReferId string `json:"refer_id,omitempty" xml:"refer_id,omitempty"`
	// 三方账号名
	ReferAccount string `json:"refer_account,omitempty" xml:"refer_account,omitempty"`
	// 站点语言ZH_CN，EN
	SiteLanguage string `json:"site_language,omitempty" xml:"site_language,omitempty"`
	// 状态 0表示正常
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 账号类型
	AccountType string `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// 是否是大陆Y表示是，N不是
	Mainland string `json:"mainland,omitempty" xml:"mainland,omitempty"`
	// 是否可用
	Available string `json:"available,omitempty" xml:"available,omitempty"`
	// 是否在职
	HrStatus string `json:"hr_status,omitempty" xml:"hr_status,omitempty"`
	// 昵称
	NickNameCn string `json:"nick_name_cn,omitempty" xml:"nick_name_cn,omitempty"`
	// 英文全名
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 中文全名
	NameCn string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
	// 中文全名
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// 英文全名
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// 工号
	EmpId string `json:"emp_id,omitempty" xml:"emp_id,omitempty"`
	// 头像
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 账号ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 租户ID
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// 是否激活
	ActiveLevel int64 `json:"active_level,omitempty" xml:"active_level,omitempty"`
}

var poolAlibabaMoziBucAccountPageallT = sync.Pool{
	New: func() any {
		return new(AlibabaMoziBucAccountPageallT)
	},
}

// GetAlibabaMoziBucAccountPageallT() 从对象池中获取AlibabaMoziBucAccountPageallT
func GetAlibabaMoziBucAccountPageallT() *AlibabaMoziBucAccountPageallT {
	return poolAlibabaMoziBucAccountPageallT.Get().(*AlibabaMoziBucAccountPageallT)
}

// ReleaseAlibabaMoziBucAccountPageallT 释放AlibabaMoziBucAccountPageallT
func ReleaseAlibabaMoziBucAccountPageallT(v *AlibabaMoziBucAccountPageallT) {
	v.Account = ""
	v.Namespace = ""
	v.ReferId = ""
	v.ReferAccount = ""
	v.SiteLanguage = ""
	v.Status = ""
	v.AccountType = ""
	v.Mainland = ""
	v.Available = ""
	v.HrStatus = ""
	v.NickNameCn = ""
	v.NameEn = ""
	v.NameCn = ""
	v.LastName = ""
	v.FirstName = ""
	v.EmpId = ""
	v.Avatar = ""
	v.Id = 0
	v.TenantId = 0
	v.ActiveLevel = 0
	poolAlibabaMoziBucAccountPageallT.Put(v)
}
