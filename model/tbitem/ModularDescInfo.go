package tbitem

import (
	"sync"
)

// ModularDescInfo 结构体
type ModularDescInfo struct {
	// 运营定义的该商品所属类目的类目级别模块信息列表，列表顺序即为模块顺序。
	ItmDscModules []ItemDescModule `json:"itm_dsc_modules,omitempty" xml:"itm_dsc_modules>item_desc_module,omitempty"`
	// 旧描述与新描述共存时间（切换时间）。
	DeadLine string `json:"dead_line,omitempty" xml:"dead_line,omitempty"`
	// 用户自定义模块数量最大值限制。类目级别模块+用户级别模块必须小于&lt;20
	UsrDefMaxNum int64 `json:"usr_def_max_num,omitempty" xml:"usr_def_max_num,omitempty"`
	// 默认值为false，如果此字段为true，则卖家上传的模块化的描述信息可以自定义排序。
	UserOrder bool `json:"user_order,omitempty" xml:"user_order,omitempty"`
}

var poolModularDescInfo = sync.Pool{
	New: func() any {
		return new(ModularDescInfo)
	},
}

// GetModularDescInfo() 从对象池中获取ModularDescInfo
func GetModularDescInfo() *ModularDescInfo {
	return poolModularDescInfo.Get().(*ModularDescInfo)
}

// ReleaseModularDescInfo 释放ModularDescInfo
func ReleaseModularDescInfo(v *ModularDescInfo) {
	v.ItmDscModules = v.ItmDscModules[:0]
	v.DeadLine = ""
	v.UsrDefMaxNum = 0
	v.UserOrder = false
	poolModularDescInfo.Put(v)
}
