package tbtrade

import (
	"sync"
)

// MonitorUrlTopDto 结构体
type MonitorUrlTopDto struct {
	// 曝光监测
	PvMonitorUrl string `json:"pv_monitor_url,omitempty" xml:"pv_monitor_url,omitempty"`
	// 点击监测
	ClickMonitorUrl string `json:"click_monitor_url,omitempty" xml:"click_monitor_url,omitempty"`
	// 落地页同步点击链接
	ClickUrl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 唤端链接
	DeepLink string `json:"deep_link,omitempty" xml:"deep_link,omitempty"`
	// 唤端链接-for ios
	UniversalLink string `json:"universal_link,omitempty" xml:"universal_link,omitempty"`
	// 监测唯一id
	MonitorUniqueId string `json:"monitor_unique_id,omitempty" xml:"monitor_unique_id,omitempty"`
}

var poolMonitorUrlTopDto = sync.Pool{
	New: func() any {
		return new(MonitorUrlTopDto)
	},
}

// GetMonitorUrlTopDto() 从对象池中获取MonitorUrlTopDto
func GetMonitorUrlTopDto() *MonitorUrlTopDto {
	return poolMonitorUrlTopDto.Get().(*MonitorUrlTopDto)
}

// ReleaseMonitorUrlTopDto 释放MonitorUrlTopDto
func ReleaseMonitorUrlTopDto(v *MonitorUrlTopDto) {
	v.PvMonitorUrl = ""
	v.ClickMonitorUrl = ""
	v.ClickUrl = ""
	v.DeepLink = ""
	v.UniversalLink = ""
	v.MonitorUniqueId = ""
	poolMonitorUrlTopDto.Put(v)
}
