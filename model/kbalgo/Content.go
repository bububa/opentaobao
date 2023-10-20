package kbalgo

import (
	"sync"
)

// Content 结构体
type Content struct {
	// 到家信息
	HomeProduct *HomeProduct `json:"home_product,omitempty" xml:"home_product,omitempty"`
	// Poi
	Poi *Poi `json:"poi,omitempty" xml:"poi,omitempty"`
	// 到店信息
	ShopProduct *ShopProduct `json:"shop_product,omitempty" xml:"shop_product,omitempty"`
}

var poolContent = sync.Pool{
	New: func() any {
		return new(Content)
	},
}

// GetContent() 从对象池中获取Content
func GetContent() *Content {
	return poolContent.Get().(*Content)
}

// ReleaseContent 释放Content
func ReleaseContent(v *Content) {
	v.HomeProduct = nil
	v.Poi = nil
	v.ShopProduct = nil
	poolContent.Put(v)
}
