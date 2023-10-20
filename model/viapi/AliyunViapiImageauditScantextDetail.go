package viapi

import (
	"sync"
)

// AliyunViapiImageauditScantextDetail 结构体
type AliyunViapiImageauditScantextDetail struct {
	// 命中该风险的上下文信息
	Contexts []Context `json:"contexts,omitempty" xml:"contexts>context,omitempty"`
	// 文本命中风险的分类
	Label string `json:"label,omitempty" xml:"label,omitempty"`
}

var poolAliyunViapiImageauditScantextDetail = sync.Pool{
	New: func() any {
		return new(AliyunViapiImageauditScantextDetail)
	},
}

// GetAliyunViapiImageauditScantextDetail() 从对象池中获取AliyunViapiImageauditScantextDetail
func GetAliyunViapiImageauditScantextDetail() *AliyunViapiImageauditScantextDetail {
	return poolAliyunViapiImageauditScantextDetail.Get().(*AliyunViapiImageauditScantextDetail)
}

// ReleaseAliyunViapiImageauditScantextDetail 释放AliyunViapiImageauditScantextDetail
func ReleaseAliyunViapiImageauditScantextDetail(v *AliyunViapiImageauditScantextDetail) {
	v.Contexts = v.Contexts[:0]
	v.Label = ""
	poolAliyunViapiImageauditScantextDetail.Put(v)
}
