package wlb

import (
	"sync"
)

// CustomData 结构体
type CustomData struct {
	//  自定义区数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	//  自定义区链接
	TemplateUrl string `json:"template_url,omitempty" xml:"template_url,omitempty"`
}

var poolCustomData = sync.Pool{
	New: func() any {
		return new(CustomData)
	},
}

// GetCustomData() 从对象池中获取CustomData
func GetCustomData() *CustomData {
	return poolCustomData.Get().(*CustomData)
}

// ReleaseCustomData 释放CustomData
func ReleaseCustomData(v *CustomData) {
	v.Data = ""
	v.TemplateUrl = ""
	poolCustomData.Put(v)
}
