package miniapp

import (
	"sync"
)

// Button 结构体
type Button struct {
	// 文案
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 跳转链接
	TargetUrl string `json:"target_url,omitempty" xml:"target_url,omitempty"`
}

var poolButton = sync.Pool{
	New: func() any {
		return new(Button)
	},
}

// GetButton() 从对象池中获取Button
func GetButton() *Button {
	return poolButton.Get().(*Button)
}

// ReleaseButton 释放Button
func ReleaseButton(v *Button) {
	v.Title = ""
	v.TargetUrl = ""
	poolButton.Put(v)
}
