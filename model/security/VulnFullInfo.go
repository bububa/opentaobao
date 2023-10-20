package security

import (
	"sync"
)

// VulnFullInfo 结构体
type VulnFullInfo struct {
	// 漏洞详情列表(任务完成时才返回)
	VulnDetails []VulnDetail `json:"vuln_details,omitempty" xml:"vuln_details>vuln_detail,omitempty"`
	// 漏洞任务错误码 0-成功 其他-错误
	TaskErrorCode string `json:"task_error_code,omitempty" xml:"task_error_code,omitempty"`
	// 漏洞任务错误信息 success-成功 其他-错误
	TaskErrorMsg string `json:"task_error_msg,omitempty" xml:"task_error_msg,omitempty"`
	// 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 漏洞数量信息(任务完成时才返回)
	VulnCount *VulnCount `json:"vuln_count,omitempty" xml:"vuln_count,omitempty"`
}

var poolVulnFullInfo = sync.Pool{
	New: func() any {
		return new(VulnFullInfo)
	},
}

// GetVulnFullInfo() 从对象池中获取VulnFullInfo
func GetVulnFullInfo() *VulnFullInfo {
	return poolVulnFullInfo.Get().(*VulnFullInfo)
}

// ReleaseVulnFullInfo 释放VulnFullInfo
func ReleaseVulnFullInfo(v *VulnFullInfo) {
	v.VulnDetails = v.VulnDetails[:0]
	v.TaskErrorCode = ""
	v.TaskErrorMsg = ""
	v.Status = 0
	v.VulnCount = nil
	poolVulnFullInfo.Put(v)
}
