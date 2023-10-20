package aedropshiper

import (
	"sync"
)

// AeopAeMultimedia 结构体
type AeopAeMultimedia struct {
	// 多媒体信息。
	AeopAEVideos []AeopAeVideo `json:"aeop_a_e_videos,omitempty" xml:"aeop_a_e_videos>aeop_ae_video,omitempty"`
}

var poolAeopAeMultimedia = sync.Pool{
	New: func() any {
		return new(AeopAeMultimedia)
	},
}

// GetAeopAeMultimedia() 从对象池中获取AeopAeMultimedia
func GetAeopAeMultimedia() *AeopAeMultimedia {
	return poolAeopAeMultimedia.Get().(*AeopAeMultimedia)
}

// ReleaseAeopAeMultimedia 释放AeopAeMultimedia
func ReleaseAeopAeMultimedia(v *AeopAeMultimedia) {
	v.AeopAEVideos = v.AeopAEVideos[:0]
	poolAeopAeMultimedia.Put(v)
}
