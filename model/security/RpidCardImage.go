package security

import (
	"sync"
)

// RpidCardImage 结构体
type RpidCardImage struct {
	// backImageUrl
	BackImageUrl string `json:"back_image_url,omitempty" xml:"back_image_url,omitempty"`
	// cardType
	CardType string `json:"card_type,omitempty" xml:"card_type,omitempty"`
	// frontImageUrl
	FrontImageUrl string `json:"front_image_url,omitempty" xml:"front_image_url,omitempty"`
}

var poolRpidCardImage = sync.Pool{
	New: func() any {
		return new(RpidCardImage)
	},
}

// GetRpidCardImage() 从对象池中获取RpidCardImage
func GetRpidCardImage() *RpidCardImage {
	return poolRpidCardImage.Get().(*RpidCardImage)
}

// ReleaseRpidCardImage 释放RpidCardImage
func ReleaseRpidCardImage(v *RpidCardImage) {
	v.BackImageUrl = ""
	v.CardType = ""
	v.FrontImageUrl = ""
	poolRpidCardImage.Put(v)
}
