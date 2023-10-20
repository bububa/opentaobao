package alicom

import (
	"sync"
)

// ActivityGiftInfos 结构体
type ActivityGiftInfos struct {
	// giftId
	GiftId string `json:"gift_id,omitempty" xml:"gift_id,omitempty"`
	// giftName
	GiftName string `json:"gift_name,omitempty" xml:"gift_name,omitempty"`
}

var poolActivityGiftInfos = sync.Pool{
	New: func() any {
		return new(ActivityGiftInfos)
	},
}

// GetActivityGiftInfos() 从对象池中获取ActivityGiftInfos
func GetActivityGiftInfos() *ActivityGiftInfos {
	return poolActivityGiftInfos.Get().(*ActivityGiftInfos)
}

// ReleaseActivityGiftInfos 释放ActivityGiftInfos
func ReleaseActivityGiftInfos(v *ActivityGiftInfos) {
	v.GiftId = ""
	v.GiftName = ""
	poolActivityGiftInfos.Put(v)
}
