package miniapp

import (
	"sync"
)

// TaobaoCoinAwardDeliveryResult 结构体
type TaobaoCoinAwardDeliveryResult struct {
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 副标题
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 按钮素材
	Button *Button `json:"button,omitempty" xml:"button,omitempty"`
	// 图片素材
	Image *Image `json:"image,omitempty" xml:"image,omitempty"`
}

var poolTaobaoCoinAwardDeliveryResult = sync.Pool{
	New: func() any {
		return new(TaobaoCoinAwardDeliveryResult)
	},
}

// GetTaobaoCoinAwardDeliveryResult() 从对象池中获取TaobaoCoinAwardDeliveryResult
func GetTaobaoCoinAwardDeliveryResult() *TaobaoCoinAwardDeliveryResult {
	return poolTaobaoCoinAwardDeliveryResult.Get().(*TaobaoCoinAwardDeliveryResult)
}

// ReleaseTaobaoCoinAwardDeliveryResult 释放TaobaoCoinAwardDeliveryResult
func ReleaseTaobaoCoinAwardDeliveryResult(v *TaobaoCoinAwardDeliveryResult) {
	v.Title = ""
	v.Desc = ""
	v.Button = nil
	v.Image = nil
	poolTaobaoCoinAwardDeliveryResult.Put(v)
}
