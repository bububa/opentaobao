package tbitem

import (
	"sync"
)

// ItemDescModule 结构体
type ItemDescModule struct {
	// 一个List&lt;String&gt;的Json串，里面是模块引导提示，不超过3条（isv提交时可忽略不传）
	Intros string `json:"intros,omitempty" xml:"intros,omitempty"`
	// 淘宝图片空间的地址链接，示例模板，最多不超过三个（isv提交时可忽略不传）
	TplUrls string `json:"tpl_urls,omitempty" xml:"tpl_urls,omitempty"`
	// 模块名称
	ModuleName string `json:"module_name,omitempty" xml:"module_name,omitempty"`
	// cat_mod:表示是运营设置的类目维度模块，usr_mod表示的是商家自定义模块。
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 描述具体内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 模块id ,(注意:用户自定义模块不用填此项。)
	ModuleId int64 `json:"module_id,omitempty" xml:"module_id,omitempty"`
	// 是否必填 （isv提交时可忽略不传）
	Required bool `json:"required,omitempty" xml:"required,omitempty"`
}

var poolItemDescModule = sync.Pool{
	New: func() any {
		return new(ItemDescModule)
	},
}

// GetItemDescModule() 从对象池中获取ItemDescModule
func GetItemDescModule() *ItemDescModule {
	return poolItemDescModule.Get().(*ItemDescModule)
}

// ReleaseItemDescModule 释放ItemDescModule
func ReleaseItemDescModule(v *ItemDescModule) {
	v.Intros = ""
	v.TplUrls = ""
	v.ModuleName = ""
	v.Type = ""
	v.Content = ""
	v.ModuleId = 0
	v.Required = false
	poolItemDescModule.Put(v)
}
