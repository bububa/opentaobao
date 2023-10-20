package aesolution

import (
	"sync"
)

// GlobalAeopAeMultimedia 结构体
type GlobalAeopAeMultimedia struct {
	// video information
	AeopAEVideos []GlobalAeopAeVideo `json:"aeop_a_e_videos,omitempty" xml:"aeop_a_e_videos>global_aeop_ae_video,omitempty"`
}

var poolGlobalAeopAeMultimedia = sync.Pool{
	New: func() any {
		return new(GlobalAeopAeMultimedia)
	},
}

// GetGlobalAeopAeMultimedia() 从对象池中获取GlobalAeopAeMultimedia
func GetGlobalAeopAeMultimedia() *GlobalAeopAeMultimedia {
	return poolGlobalAeopAeMultimedia.Get().(*GlobalAeopAeMultimedia)
}

// ReleaseGlobalAeopAeMultimedia 释放GlobalAeopAeMultimedia
func ReleaseGlobalAeopAeMultimedia(v *GlobalAeopAeMultimedia) {
	v.AeopAEVideos = v.AeopAEVideos[:0]
	poolGlobalAeopAeMultimedia.Put(v)
}
