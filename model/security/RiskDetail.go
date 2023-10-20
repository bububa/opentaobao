package security

import (
	"sync"
)

// RiskDetail 结构体
type RiskDetail struct {
	// 仿冒应用信息
	FakeInfo *FakeAppFullInfo `json:"fake_info,omitempty" xml:"fake_info,omitempty"`
	// 恶意代码信息
	MalwareInfo *MalwareFullInfo `json:"malware_info,omitempty" xml:"malware_info,omitempty"`
	// 插件信息
	PluginInfo *PluginFullInfo `json:"plugin_info,omitempty" xml:"plugin_info,omitempty"`
	// 任务状态: 1-已完成,2-处理中,3-处理失败,4-处理超时
	TaskStatus int64 `json:"task_status,omitempty" xml:"task_status,omitempty"`
	// 漏洞信息
	VulnInfo *VulnFullInfo `json:"vuln_info,omitempty" xml:"vuln_info,omitempty"`
}

var poolRiskDetail = sync.Pool{
	New: func() any {
		return new(RiskDetail)
	},
}

// GetRiskDetail() 从对象池中获取RiskDetail
func GetRiskDetail() *RiskDetail {
	return poolRiskDetail.Get().(*RiskDetail)
}

// ReleaseRiskDetail 释放RiskDetail
func ReleaseRiskDetail(v *RiskDetail) {
	v.FakeInfo = nil
	v.MalwareInfo = nil
	v.PluginInfo = nil
	v.TaskStatus = 0
	v.VulnInfo = nil
	poolRiskDetail.Put(v)
}
