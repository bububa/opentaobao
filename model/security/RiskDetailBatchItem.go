package security

import (
	"sync"
)

// RiskDetailBatchItem 结构体
type RiskDetailBatchItem struct {
	// app的md5
	AppIdentity string `json:"app_identity,omitempty" xml:"app_identity,omitempty"`
	// 恶意代码检测结果
	MalwareInfo *MalwareFullInfo `json:"malware_info,omitempty" xml:"malware_info,omitempty"`
	// 插件信息
	PluginInfo *PluginFullInfo `json:"plugin_info,omitempty" xml:"plugin_info,omitempty"`
	// 漏洞列表(任务完成时才返回)
	VulnInfo *VulnFullInfo `json:"vuln_info,omitempty" xml:"vuln_info,omitempty"`
}

var poolRiskDetailBatchItem = sync.Pool{
	New: func() any {
		return new(RiskDetailBatchItem)
	},
}

// GetRiskDetailBatchItem() 从对象池中获取RiskDetailBatchItem
func GetRiskDetailBatchItem() *RiskDetailBatchItem {
	return poolRiskDetailBatchItem.Get().(*RiskDetailBatchItem)
}

// ReleaseRiskDetailBatchItem 释放RiskDetailBatchItem
func ReleaseRiskDetailBatchItem(v *RiskDetailBatchItem) {
	v.AppIdentity = ""
	v.MalwareInfo = nil
	v.PluginInfo = nil
	v.VulnInfo = nil
	poolRiskDetailBatchItem.Put(v)
}
