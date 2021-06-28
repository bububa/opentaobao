package security

// VulnCount 
/* model for simplify = false
type VulnCount struct {

    // 漏洞总数量
    
    Total   int64 `json:"total,omitempty"`
    

    // 高风险漏洞数量
    
    HighLevel   int64 `json:"high_level,omitempty"`
    

    // 中风险漏洞数量
    
    MidLevel   int64 `json:"mid_level,omitempty"`
    

    // 低风险漏洞数量
    
    LowLevel   int64 `json:"low_level,omitempty"`
    

    // 安全红线漏洞数量
    
    RedLine   int64 `json:"red_line,omitempty"`
    

}
*/

// VulnCount 
type VulnCount struct {

    // 漏洞总数量
    Total   int64 `json:"total,omitempty"`

    // 高风险漏洞数量
    HighLevel   int64 `json:"high_level,omitempty"`

    // 中风险漏洞数量
    MidLevel   int64 `json:"mid_level,omitempty"`

    // 低风险漏洞数量
    LowLevel   int64 `json:"low_level,omitempty"`

    // 安全红线漏洞数量
    RedLine   int64 `json:"red_line,omitempty"`

}
