package viapi

import (
	"sync"
)

// AliyunViapiImageauditScantextData 结构体
type AliyunViapiImageauditScantextData struct {
	// 检测结果各个子元素
	Elements []Element `json:"elements,omitempty" xml:"elements>element,omitempty"`
}

var poolAliyunViapiImageauditScantextData = sync.Pool{
	New: func() any {
		return new(AliyunViapiImageauditScantextData)
	},
}

// GetAliyunViapiImageauditScantextData() 从对象池中获取AliyunViapiImageauditScantextData
func GetAliyunViapiImageauditScantextData() *AliyunViapiImageauditScantextData {
	return poolAliyunViapiImageauditScantextData.Get().(*AliyunViapiImageauditScantextData)
}

// ReleaseAliyunViapiImageauditScantextData 释放AliyunViapiImageauditScantextData
func ReleaseAliyunViapiImageauditScantextData(v *AliyunViapiImageauditScantextData) {
	v.Elements = v.Elements[:0]
	poolAliyunViapiImageauditScantextData.Put(v)
}
