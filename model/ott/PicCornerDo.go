package ott

import (
	"sync"
)

// PicCornerDo 结构体
type PicCornerDo struct {
	// 角标类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 角标地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 角标文案
	Text string `json:"text,omitempty" xml:"text,omitempty"`
}

var poolPicCornerDo = sync.Pool{
	New: func() any {
		return new(PicCornerDo)
	},
}

// GetPicCornerDo() 从对象池中获取PicCornerDo
func GetPicCornerDo() *PicCornerDo {
	return poolPicCornerDo.Get().(*PicCornerDo)
}

// ReleasePicCornerDo 释放PicCornerDo
func ReleasePicCornerDo(v *PicCornerDo) {
	v.Type = ""
	v.Url = ""
	v.Text = ""
	poolPicCornerDo.Put(v)
}
