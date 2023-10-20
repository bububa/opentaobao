package campus

import (
	"sync"
)

// DeviceDataApiQuery 结构体
type DeviceDataApiQuery struct {
	// 启始时间(防止接口超时,建议不要传入时间跨度过大,如查询一个月内的数据)
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 参数点code(详细请查阅‘平台技术’下‘设备详细信息开发文档’。)
	PropertyCode string `json:"property_code,omitempty" xml:"property_code,omitempty"`
	// 终止时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 设备id
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 分页大小(最大500)
	Limit int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}

var poolDeviceDataApiQuery = sync.Pool{
	New: func() any {
		return new(DeviceDataApiQuery)
	},
}

// GetDeviceDataApiQuery() 从对象池中获取DeviceDataApiQuery
func GetDeviceDataApiQuery() *DeviceDataApiQuery {
	return poolDeviceDataApiQuery.Get().(*DeviceDataApiQuery)
}

// ReleaseDeviceDataApiQuery 释放DeviceDataApiQuery
func ReleaseDeviceDataApiQuery(v *DeviceDataApiQuery) {
	v.StartDate = ""
	v.PropertyCode = ""
	v.EndDate = ""
	v.Uuid = ""
	v.Limit = 0
	v.CurrentPage = 0
	poolDeviceDataApiQuery.Put(v)
}
