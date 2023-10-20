package idleisv

import (
	"sync"
)

// IdleItemApiAutoRechargeDo 结构体
type IdleItemApiAutoRechargeDo struct {
	// 直充模板code
	TemplateCode string `json:"template_code,omitempty" xml:"template_code,omitempty"`
	// 模版的额外参数
	TemplateExtraInfo string `json:"template_extra_info,omitempty" xml:"template_extra_info,omitempty"`
	// 直充动作，add:添加，remove:移除
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 当前服务商是否拥有直充编辑权限
	Owner bool `json:"owner,omitempty" xml:"owner,omitempty"`
}

var poolIdleItemApiAutoRechargeDo = sync.Pool{
	New: func() any {
		return new(IdleItemApiAutoRechargeDo)
	},
}

// GetIdleItemApiAutoRechargeDo() 从对象池中获取IdleItemApiAutoRechargeDo
func GetIdleItemApiAutoRechargeDo() *IdleItemApiAutoRechargeDo {
	return poolIdleItemApiAutoRechargeDo.Get().(*IdleItemApiAutoRechargeDo)
}

// ReleaseIdleItemApiAutoRechargeDo 释放IdleItemApiAutoRechargeDo
func ReleaseIdleItemApiAutoRechargeDo(v *IdleItemApiAutoRechargeDo) {
	v.TemplateCode = ""
	v.TemplateExtraInfo = ""
	v.Action = ""
	v.Owner = false
	poolIdleItemApiAutoRechargeDo.Put(v)
}
