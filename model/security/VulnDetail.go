package security

import (
	"sync"
)

// VulnDetail 结构体
type VulnDetail struct {
	// 漏洞位置
	Locations []string `json:"locations,omitempty" xml:"locations>string,omitempty"`
	// 漏洞风险级别: 高， 中， 低
	Level string `json:"level,omitempty" xml:"level,omitempty"`
	// 漏洞名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 漏洞修复建议
	Recommendation string `json:"recommendation,omitempty" xml:"recommendation,omitempty"`
	// 漏洞详细说明参考链接
	ReferenctLink string `json:"referenct_link,omitempty" xml:"referenct_link,omitempty"`
	// 漏洞类型码
	VulnId string `json:"vuln_id,omitempty" xml:"vuln_id,omitempty"`
	// 漏洞描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 漏洞数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 是否安全红线漏洞
	RedLine bool `json:"red_line,omitempty" xml:"red_line,omitempty"`
}

var poolVulnDetail = sync.Pool{
	New: func() any {
		return new(VulnDetail)
	},
}

// GetVulnDetail() 从对象池中获取VulnDetail
func GetVulnDetail() *VulnDetail {
	return poolVulnDetail.Get().(*VulnDetail)
}

// ReleaseVulnDetail 释放VulnDetail
func ReleaseVulnDetail(v *VulnDetail) {
	v.Locations = v.Locations[:0]
	v.Level = ""
	v.Name = ""
	v.Recommendation = ""
	v.ReferenctLink = ""
	v.VulnId = ""
	v.Description = ""
	v.Count = 0
	v.RedLine = false
	poolVulnDetail.Put(v)
}
