package rail

import (
	"sync"
)

// RailStationRs 结构体
type RailStationRs struct {
	// 车站图片url，多个;号分隔
	Image string `json:"image,omitempty" xml:"image,omitempty"`
	// 车站信息，多行;号分隔
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 车站地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 车站中文名称
	CnName string `json:"cn_name,omitempty" xml:"cn_name,omitempty"`
	// 车站英文名称
	EnName string `json:"en_name,omitempty" xml:"en_name,omitempty"`
	// 所属城市名称
	DivisionName string `json:"division_name,omitempty" xml:"division_name,omitempty"`
	// 车站原始名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 车站编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 所属城市id
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
}

var poolRailStationRs = sync.Pool{
	New: func() any {
		return new(RailStationRs)
	},
}

// GetRailStationRs() 从对象池中获取RailStationRs
func GetRailStationRs() *RailStationRs {
	return poolRailStationRs.Get().(*RailStationRs)
}

// ReleaseRailStationRs 释放RailStationRs
func ReleaseRailStationRs(v *RailStationRs) {
	v.Image = ""
	v.Detail = ""
	v.Address = ""
	v.Latitude = ""
	v.Longitude = ""
	v.CnName = ""
	v.EnName = ""
	v.DivisionName = ""
	v.Name = ""
	v.Code = ""
	v.DivisionId = 0
	poolRailStationRs.Put(v)
}
