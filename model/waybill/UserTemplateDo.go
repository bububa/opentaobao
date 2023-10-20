package waybill

import (
	"sync"
)

// UserTemplateDo 结构体
type UserTemplateDo struct {
	// keys
	Keys []KeyResult `json:"keys,omitempty" xml:"keys>key_result,omitempty"`
	// 用户使用模板的url
	UserStdTemplateUrl string `json:"user_std_template_url,omitempty" xml:"user_std_template_url,omitempty"`
	// 用户使用模板名称
	UserStdTemplateName string `json:"user_std_template_name,omitempty" xml:"user_std_template_name,omitempty"`
	// 品牌 code
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 用户使用模板的id
	UserStdTemplateId int64 `json:"user_std_template_id,omitempty" xml:"user_std_template_id,omitempty"`
}

var poolUserTemplateDo = sync.Pool{
	New: func() any {
		return new(UserTemplateDo)
	},
}

// GetUserTemplateDo() 从对象池中获取UserTemplateDo
func GetUserTemplateDo() *UserTemplateDo {
	return poolUserTemplateDo.Get().(*UserTemplateDo)
}

// ReleaseUserTemplateDo 释放UserTemplateDo
func ReleaseUserTemplateDo(v *UserTemplateDo) {
	v.Keys = v.Keys[:0]
	v.UserStdTemplateUrl = ""
	v.UserStdTemplateName = ""
	v.BrandCode = ""
	v.UserStdTemplateId = 0
	poolUserTemplateDo.Put(v)
}
