package tvupadmin

import (
	"sync"
)

// AdvertScheduleDo 结构体
type AdvertScheduleDo struct {
	// 设备型号
	DeviceModel string `json:"device_model,omitempty" xml:"device_model,omitempty"`
	// UUID
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 开始时间
	Start string `json:"start,omitempty" xml:"start,omitempty"`
	// 结束时间
	End string `json:"end,omitempty" xml:"end,omitempty"`
	// 查询类型
	Range int64 `json:"range,omitempty" xml:"range,omitempty"`
	// 播控ID
	BcpId int64 `json:"bcp_id,omitempty" xml:"bcp_id,omitempty"`
	// 广告类型
	SiteType int64 `json:"site_type,omitempty" xml:"site_type,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolAdvertScheduleDo = sync.Pool{
	New: func() any {
		return new(AdvertScheduleDo)
	},
}

// GetAdvertScheduleDo() 从对象池中获取AdvertScheduleDo
func GetAdvertScheduleDo() *AdvertScheduleDo {
	return poolAdvertScheduleDo.Get().(*AdvertScheduleDo)
}

// ReleaseAdvertScheduleDo 释放AdvertScheduleDo
func ReleaseAdvertScheduleDo(v *AdvertScheduleDo) {
	v.DeviceModel = ""
	v.Uuid = ""
	v.Start = ""
	v.End = ""
	v.Range = 0
	v.BcpId = 0
	v.SiteType = 0
	v.Status = 0
	poolAdvertScheduleDo.Put(v)
}
