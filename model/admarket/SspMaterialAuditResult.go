package admarket

import (
	"sync"
)

// SspMaterialAuditResult 结构体
type SspMaterialAuditResult struct {
	// 排除设备
	ExcludeDevices []ExcludeDevice `json:"exclude_devices,omitempty" xml:"exclude_devices>exclude_device,omitempty"`
	// 原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 审核结果
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 创意id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
}

var poolSspMaterialAuditResult = sync.Pool{
	New: func() any {
		return new(SspMaterialAuditResult)
	},
}

// GetSspMaterialAuditResult() 从对象池中获取SspMaterialAuditResult
func GetSspMaterialAuditResult() *SspMaterialAuditResult {
	return poolSspMaterialAuditResult.Get().(*SspMaterialAuditResult)
}

// ReleaseSspMaterialAuditResult 释放SspMaterialAuditResult
func ReleaseSspMaterialAuditResult(v *SspMaterialAuditResult) {
	v.ExcludeDevices = v.ExcludeDevices[:0]
	v.Reason = ""
	v.Status = ""
	v.Channel = ""
	v.MaterialId = 0
	poolSspMaterialAuditResult.Put(v)
}
