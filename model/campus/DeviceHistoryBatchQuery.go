package campus

import (
	"sync"
)

// DeviceHistoryBatchQuery 结构体
type DeviceHistoryBatchQuery struct {
	// 设备列表
	DeviceUuidList []string `json:"device_uuid_list,omitempty" xml:"device_uuid_list>string,omitempty"`
	// 历史数据查询开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 历史数据查询结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 设备模板code
	KindCode string `json:"kind_code,omitempty" xml:"kind_code,omitempty"`
	// 设备属性code
	PropertyCode string `json:"property_code,omitempty" xml:"property_code,omitempty"`
	// 设备所属园区id
	CampusId string `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 历史数据查询间隔（min）
	IntervalMinutes int64 `json:"interval_minutes,omitempty" xml:"interval_minutes,omitempty"`
}

var poolDeviceHistoryBatchQuery = sync.Pool{
	New: func() any {
		return new(DeviceHistoryBatchQuery)
	},
}

// GetDeviceHistoryBatchQuery() 从对象池中获取DeviceHistoryBatchQuery
func GetDeviceHistoryBatchQuery() *DeviceHistoryBatchQuery {
	return poolDeviceHistoryBatchQuery.Get().(*DeviceHistoryBatchQuery)
}

// ReleaseDeviceHistoryBatchQuery 释放DeviceHistoryBatchQuery
func ReleaseDeviceHistoryBatchQuery(v *DeviceHistoryBatchQuery) {
	v.DeviceUuidList = v.DeviceUuidList[:0]
	v.StartDate = ""
	v.EndDate = ""
	v.KindCode = ""
	v.PropertyCode = ""
	v.CampusId = ""
	v.IntervalMinutes = 0
	poolDeviceHistoryBatchQuery.Put(v)
}
