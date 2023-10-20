package travel

import (
	"sync"
)

// FreeTourTrafficInfo 结构体
type FreeTourTrafficInfo struct {
	// 参考班次号，飞机需要填航班号，火车需要填车次号，汽车和船可不填
	TrafficNo string `json:"traffic_no,omitempty" xml:"traffic_no,omitempty"`
	// 交通公司名，飞机选填
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 飞机机型，飞机选填
	PlaneType string `json:"plane_type,omitempty" xml:"plane_type,omitempty"`
	// 出发城市
	Departure string `json:"departure,omitempty" xml:"departure,omitempty"`
	// 到达城市
	Destination string `json:"destination,omitempty" xml:"destination,omitempty"`
	// 出发时间，当地时间HH:mm
	DepartureTime string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// 到达时间，当地时间HH:mm
	ArrivalTime string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// 交通说明，针对交通类型是汽车，轮船和其他
	TrafficDesc string `json:"traffic_desc,omitempty" xml:"traffic_desc,omitempty"`
	// 经停城市
	StopCity string `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	// 第几天
	Day int64 `json:"day,omitempty" xml:"day,omitempty"`
	// 是否直飞，飞机选填，1-直飞；0-不是直飞
	NonStop int64 `json:"non_stop,omitempty" xml:"non_stop,omitempty"`
	// 第几组交通信息，请必传，不传可能导致商品详情页交通信息不展示。每一组交通信息包含一组去程交通和返程交通，当在页&gt; 面上点击【添加交通信息】按钮后，就会出现第二组交通信息，第一组交&gt;通信息group=1，第二组交通信息group取值为2，以此类推
	Group int64 `json:"group,omitempty" xml:"group,omitempty"`
	// 是否经停
	StopOver bool `json:"stop_over,omitempty" xml:"stop_over,omitempty"`
}

var poolFreeTourTrafficInfo = sync.Pool{
	New: func() any {
		return new(FreeTourTrafficInfo)
	},
}

// GetFreeTourTrafficInfo() 从对象池中获取FreeTourTrafficInfo
func GetFreeTourTrafficInfo() *FreeTourTrafficInfo {
	return poolFreeTourTrafficInfo.Get().(*FreeTourTrafficInfo)
}

// ReleaseFreeTourTrafficInfo 释放FreeTourTrafficInfo
func ReleaseFreeTourTrafficInfo(v *FreeTourTrafficInfo) {
	v.TrafficNo = ""
	v.Vendor = ""
	v.PlaneType = ""
	v.Departure = ""
	v.Destination = ""
	v.DepartureTime = ""
	v.ArrivalTime = ""
	v.TrafficDesc = ""
	v.StopCity = ""
	v.Day = 0
	v.NonStop = 0
	v.Group = 0
	v.StopOver = false
	poolFreeTourTrafficInfo.Put(v)
}
