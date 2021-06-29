package security

// VulnFullInfo 
type VulnFullInfo struct {
    // 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 漏洞数量信息(任务完成时才返回)
    VulnCount   *VulnCount `json:"vuln_count,omitempty" xml:"vuln_count,omitempty"`
    // 漏洞详情列表(任务完成时才返回)
    VulnDetails   []VulnDetail `json:"vuln_details,omitempty" xml:"vuln_details>vuln_detail,omitempty"`
    // 漏洞任务错误码 0-成功 其他-错误
    TaskErrorCode   string `json:"task_error_code,omitempty" xml:"task_error_code,omitempty"`
    // 漏洞任务错误信息 success-成功 其他-错误
    TaskErrorMsg   string `json:"task_error_msg,omitempty" xml:"task_error_msg,omitempty"`
}
