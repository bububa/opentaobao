package eleenterpriseordernew

import (
	"sync"
)

// TrackingInfoDto 结构体
type TrackingInfoDto struct {
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
}

var poolTrackingInfoDto = sync.Pool{
	New: func() any {
		return new(TrackingInfoDto)
	},
}

// GetTrackingInfoDto() 从对象池中获取TrackingInfoDto
func GetTrackingInfoDto() *TrackingInfoDto {
	return poolTrackingInfoDto.Get().(*TrackingInfoDto)
}

// ReleaseTrackingInfoDto 释放TrackingInfoDto
func ReleaseTrackingInfoDto(v *TrackingInfoDto) {
	v.Latitude = ""
	v.Longitude = ""
	poolTrackingInfoDto.Put(v)
}
