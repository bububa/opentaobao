package security

import (
	"sync"
)

// VulnCount 结构体
type VulnCount struct {
	// 高风险漏洞数量
	HighLevel int64 `json:"high_level,omitempty" xml:"high_level,omitempty"`
	// 低风险漏洞数量
	LowLevel int64 `json:"low_level,omitempty" xml:"low_level,omitempty"`
	// 中风险漏洞数量
	MidLevel int64 `json:"mid_level,omitempty" xml:"mid_level,omitempty"`
	// 安全红线漏洞数量
	RedLine int64 `json:"red_line,omitempty" xml:"red_line,omitempty"`
	// 漏洞总数量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolVulnCount = sync.Pool{
	New: func() any {
		return new(VulnCount)
	},
}

// GetVulnCount() 从对象池中获取VulnCount
func GetVulnCount() *VulnCount {
	return poolVulnCount.Get().(*VulnCount)
}

// ReleaseVulnCount 释放VulnCount
func ReleaseVulnCount(v *VulnCount) {
	v.HighLevel = 0
	v.LowLevel = 0
	v.MidLevel = 0
	v.RedLine = 0
	v.Total = 0
	poolVulnCount.Put(v)
}
