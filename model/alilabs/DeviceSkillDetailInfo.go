package alilabs

import (
	"sync"
)

// DeviceSkillDetailInfo 结构体
type DeviceSkillDetailInfo struct {
	// 提供商集合
	ServiceProviders []ServiceProvider `json:"service_providers,omitempty" xml:"service_providers>service_provider,omitempty"`
	// 唤醒词
	InvocationName string `json:"invocation_name,omitempty" xml:"invocation_name,omitempty"`
	// 技能名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 类别
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 技能图片地址
	IcoinImageUrl string `json:"icoin_image_url,omitempty" xml:"icoin_image_url,omitempty"`
	// 技能描述信息
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 典型例句(多个，以\t分隔)
	Sample string `json:"sample,omitempty" xml:"sample,omitempty"`
	// 技能Id
	SkillId int64 `json:"skill_id,omitempty" xml:"skill_id,omitempty"`
}

var poolDeviceSkillDetailInfo = sync.Pool{
	New: func() any {
		return new(DeviceSkillDetailInfo)
	},
}

// GetDeviceSkillDetailInfo() 从对象池中获取DeviceSkillDetailInfo
func GetDeviceSkillDetailInfo() *DeviceSkillDetailInfo {
	return poolDeviceSkillDetailInfo.Get().(*DeviceSkillDetailInfo)
}

// ReleaseDeviceSkillDetailInfo 释放DeviceSkillDetailInfo
func ReleaseDeviceSkillDetailInfo(v *DeviceSkillDetailInfo) {
	v.ServiceProviders = v.ServiceProviders[:0]
	v.InvocationName = ""
	v.Name = ""
	v.Category = ""
	v.IcoinImageUrl = ""
	v.Desc = ""
	v.Sample = ""
	v.SkillId = 0
	poolDeviceSkillDetailInfo.Put(v)
}
