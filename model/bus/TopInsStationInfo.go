package bus

import (
	"sync"
)

// TopInsStationInfo 结构体
type TopInsStationInfo struct {
	// 机具ID
	MachineId string `json:"machine_id,omitempty" xml:"machine_id,omitempty"`
	// 城市代码，数据源来自行业基础数据字典
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// isv名称
	IsvName string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// isv-id
	IsvId string `json:"isv_id,omitempty" xml:"isv_id,omitempty"`
	// 省份代码，数据源来自行业基础数据字典
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 出发车站站点ID
	StartStationId string `json:"start_station_id,omitempty" xml:"start_station_id,omitempty"`
}

var poolTopInsStationInfo = sync.Pool{
	New: func() any {
		return new(TopInsStationInfo)
	},
}

// GetTopInsStationInfo() 从对象池中获取TopInsStationInfo
func GetTopInsStationInfo() *TopInsStationInfo {
	return poolTopInsStationInfo.Get().(*TopInsStationInfo)
}

// ReleaseTopInsStationInfo 释放TopInsStationInfo
func ReleaseTopInsStationInfo(v *TopInsStationInfo) {
	v.MachineId = ""
	v.CityCode = ""
	v.IsvName = ""
	v.IsvId = ""
	v.ProvinceCode = ""
	v.StartStationId = ""
	poolTopInsStationInfo.Put(v)
}
