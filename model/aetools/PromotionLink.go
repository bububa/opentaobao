package aetools

import (
	"sync"
)

// PromotionLink 结构体
type PromotionLink struct {
	// 推广链接
	PromotionLink string `json:"promotion_link,omitempty" xml:"promotion_link,omitempty"`
	// 原始链接或者值
	SourceValue string `json:"source_value,omitempty" xml:"source_value,omitempty"`
}

var poolPromotionLink = sync.Pool{
	New: func() any {
		return new(PromotionLink)
	},
}

// GetPromotionLink() 从对象池中获取PromotionLink
func GetPromotionLink() *PromotionLink {
	return poolPromotionLink.Get().(*PromotionLink)
}

// ReleasePromotionLink 释放PromotionLink
func ReleasePromotionLink(v *PromotionLink) {
	v.PromotionLink = ""
	v.SourceValue = ""
	poolPromotionLink.Put(v)
}
