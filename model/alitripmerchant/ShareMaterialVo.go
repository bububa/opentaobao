package alitripmerchant

import (
	"sync"
)

// ShareMaterialVo 结构体
type ShareMaterialVo struct {
	// 分享的文案
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 分享的图片
	Picture string `json:"picture,omitempty" xml:"picture,omitempty"`
}

var poolShareMaterialVo = sync.Pool{
	New: func() any {
		return new(ShareMaterialVo)
	},
}

// GetShareMaterialVo() 从对象池中获取ShareMaterialVo
func GetShareMaterialVo() *ShareMaterialVo {
	return poolShareMaterialVo.Get().(*ShareMaterialVo)
}

// ReleaseShareMaterialVo 释放ShareMaterialVo
func ReleaseShareMaterialVo(v *ShareMaterialVo) {
	v.Text = ""
	v.Picture = ""
	poolShareMaterialVo.Put(v)
}
