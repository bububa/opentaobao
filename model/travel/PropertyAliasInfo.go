package travel

import (
	"sync"
)

// PropertyAliasInfo 结构体
type PropertyAliasInfo struct {
	// 销售属性的pid和vid，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 属性具体别名值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolPropertyAliasInfo = sync.Pool{
	New: func() any {
		return new(PropertyAliasInfo)
	},
}

// GetPropertyAliasInfo() 从对象池中获取PropertyAliasInfo
func GetPropertyAliasInfo() *PropertyAliasInfo {
	return poolPropertyAliasInfo.Get().(*PropertyAliasInfo)
}

// ReleasePropertyAliasInfo 释放PropertyAliasInfo
func ReleasePropertyAliasInfo(v *PropertyAliasInfo) {
	v.Properties = ""
	v.Value = ""
	poolPropertyAliasInfo.Put(v)
}
