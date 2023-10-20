package waybill

import (
	"sync"
)

// UserTemplateResult 结构体
type UserTemplateResult struct {
	// 用户使用的模板数据
	UserStdTemplates []UserTemplateDo `json:"user_std_templates,omitempty" xml:"user_std_templates>user_template_do,omitempty"`
	// cp编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
}

var poolUserTemplateResult = sync.Pool{
	New: func() any {
		return new(UserTemplateResult)
	},
}

// GetUserTemplateResult() 从对象池中获取UserTemplateResult
func GetUserTemplateResult() *UserTemplateResult {
	return poolUserTemplateResult.Get().(*UserTemplateResult)
}

// ReleaseUserTemplateResult 释放UserTemplateResult
func ReleaseUserTemplateResult(v *UserTemplateResult) {
	v.UserStdTemplates = v.UserStdTemplates[:0]
	v.CpCode = ""
	poolUserTemplateResult.Put(v)
}
