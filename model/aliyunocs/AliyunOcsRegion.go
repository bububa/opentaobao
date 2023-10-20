package aliyunocs

import (
	"sync"
)

// AliyunOcsRegion 结构体
type AliyunOcsRegion struct {
	// cn-hangzhou
	RegionId string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// &#34;B,D&#34;
	ZoneIds string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
	// 杭州
	LocalName string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
}

var poolAliyunOcsRegion = sync.Pool{
	New: func() any {
		return new(AliyunOcsRegion)
	},
}

// GetAliyunOcsRegion() 从对象池中获取AliyunOcsRegion
func GetAliyunOcsRegion() *AliyunOcsRegion {
	return poolAliyunOcsRegion.Get().(*AliyunOcsRegion)
}

// ReleaseAliyunOcsRegion 释放AliyunOcsRegion
func ReleaseAliyunOcsRegion(v *AliyunOcsRegion) {
	v.RegionId = ""
	v.ZoneIds = ""
	v.LocalName = ""
	poolAliyunOcsRegion.Put(v)
}
