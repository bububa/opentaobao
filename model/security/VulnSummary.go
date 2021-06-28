package security

// VulnSummary 
/* model for simplify = false
type VulnSummary struct {

    // 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
    
    Status   int64 `json:"status,omitempty"`
    

    // 漏洞数量信息(任务完成时才返回)
    
    VulnCount  *struct {
        VulnCount  *VulnCount `json:"vuln_count,omitempty"`
    } `json:"vuln_count,omitempty"`
    

    // 漏洞任务错误码 0-成功 其他-错误
    
    TaskErrorCode   string `json:"task_error_code,omitempty"`
    

    // 漏洞任务错误信息 succes-成功  其他-错误
    
    TaskErrorMsg   string `json:"task_error_msg,omitempty"`
    

}
*/

// VulnSummary 
type VulnSummary struct {

    // 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
    Status   int64 `json:"status,omitempty"`

    // 漏洞数量信息(任务完成时才返回)
    VulnCount   *VulnCount `json:"vuln_count,omitempty"`

    // 漏洞任务错误码 0-成功 其他-错误
    TaskErrorCode   string `json:"task_error_code,omitempty"`

    // 漏洞任务错误信息 succes-成功  其他-错误
    TaskErrorMsg   string `json:"task_error_msg,omitempty"`

}
