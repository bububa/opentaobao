package usergrowth

import (
	"sync"
)

// ExtraDto 结构体
type ExtraDto struct {
	// 拓展曝光上报链接, 可以用于榜单曝光
	ExposureUrl []string `json:"exposure_url,omitempty" xml:"exposure_url>string,omitempty"`
	// 拓展点击上报链接, 可以用于“点击更多”上报
	ClickUrl []string `json:"click_url,omitempty" xml:"click_url>string,omitempty"`
	// App跳转链接
	DeeplinkUrl string `json:"deeplink_url,omitempty" xml:"deeplink_url,omitempty"`
	// 未安装淘宝时 h5 页面跳
	H5Url string `json:"h5_url,omitempty" xml:"h5_url,omitempty"`
}

var poolExtraDto = sync.Pool{
	New: func() any {
		return new(ExtraDto)
	},
}

// GetExtraDto() 从对象池中获取ExtraDto
func GetExtraDto() *ExtraDto {
	return poolExtraDto.Get().(*ExtraDto)
}

// ReleaseExtraDto 释放ExtraDto
func ReleaseExtraDto(v *ExtraDto) {
	v.ExposureUrl = v.ExposureUrl[:0]
	v.ClickUrl = v.ClickUrl[:0]
	v.DeeplinkUrl = ""
	v.H5Url = ""
	poolExtraDto.Put(v)
}
