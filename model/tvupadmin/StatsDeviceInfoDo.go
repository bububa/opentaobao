package tvupadmin

import (
	"sync"
)

// StatsDeviceInfoDo 结构体
type StatsDeviceInfoDo struct {
	// factoryName
	FactoryName string `json:"factory_name,omitempty" xml:"factory_name,omitempty"`
	// deviceModel
	DeviceModel string `json:"device_model,omitempty" xml:"device_model,omitempty"`
	// statsDateStr
	StatsDateStr string `json:"stats_date_str,omitempty" xml:"stats_date_str,omitempty"`
	// statsDate
	StatsDate string `json:"stats_date,omitempty" xml:"stats_date,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// userActiveTotal
	UserActiveTotal int64 `json:"user_active_total,omitempty" xml:"user_active_total,omitempty"`
	// userSignDaily
	UserSignDaily int64 `json:"user_sign_daily,omitempty" xml:"user_sign_daily,omitempty"`
}

var poolStatsDeviceInfoDo = sync.Pool{
	New: func() any {
		return new(StatsDeviceInfoDo)
	},
}

// GetStatsDeviceInfoDo() 从对象池中获取StatsDeviceInfoDo
func GetStatsDeviceInfoDo() *StatsDeviceInfoDo {
	return poolStatsDeviceInfoDo.Get().(*StatsDeviceInfoDo)
}

// ReleaseStatsDeviceInfoDo 释放StatsDeviceInfoDo
func ReleaseStatsDeviceInfoDo(v *StatsDeviceInfoDo) {
	v.FactoryName = ""
	v.DeviceModel = ""
	v.StatsDateStr = ""
	v.StatsDate = ""
	v.Id = 0
	v.UserActiveTotal = 0
	v.UserSignDaily = 0
	poolStatsDeviceInfoDo.Put(v)
}
