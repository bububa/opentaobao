package iotticket

import (
	"sync"
)

// Images 结构体
type Images struct {
	// 图片类型：service-上门服务图片;cuSendProof-客户邮寄凭证;spSendProof-服务商邮寄凭证;abnormalImage-异常信息;purchaseVoucher-用户购买凭证
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 图片url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolImages = sync.Pool{
	New: func() any {
		return new(Images)
	},
}

// GetImages() 从对象池中获取Images
func GetImages() *Images {
	return poolImages.Get().(*Images)
}

// ReleaseImages 释放Images
func ReleaseImages(v *Images) {
	v.Type = ""
	v.Url = ""
	poolImages.Put(v)
}
