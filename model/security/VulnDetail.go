package security

// VulnDetail 
/* model for simplify = false
type VulnDetail struct {

    // 漏洞名称
    
    Name   string `json:"name,omitempty"`
    

    // 漏洞风险级别: 高， 中， 低
    
    Level   string `json:"level,omitempty"`
    

    // 漏洞数量
    
    Count   int64 `json:"count,omitempty"`
    

    // 漏洞位置
    
    Locations  struct {
        String  []string `json:"string,omitempty"`
    } `json:"locations,omitempty"`
    

    // 是否安全红线漏洞
    
    RedLine   bool `json:"red_line,omitempty"`
    

    // 漏洞详细说明参考链接
    
    ReferenctLink   string `json:"referenct_link,omitempty"`
    

    // 漏洞修复建议
    
    Recommendation   string `json:"recommendation,omitempty"`
    

    // 漏洞类型码
    
    VulnId   string `json:"vuln_id,omitempty"`
    

    // 漏洞描述
    
    Description   string `json:"description,omitempty"`
    

}
*/

// VulnDetail 
type VulnDetail struct {

    // 漏洞名称
    Name   string `json:"name,omitempty"`

    // 漏洞风险级别: 高， 中， 低
    Level   string `json:"level,omitempty"`

    // 漏洞数量
    Count   int64 `json:"count,omitempty"`

    // 漏洞位置
    Locations   []string `json:"locations,omitempty"`

    // 是否安全红线漏洞
    RedLine   bool `json:"red_line,omitempty"`

    // 漏洞详细说明参考链接
    ReferenctLink   string `json:"referenct_link,omitempty"`

    // 漏洞修复建议
    Recommendation   string `json:"recommendation,omitempty"`

    // 漏洞类型码
    VulnId   string `json:"vuln_id,omitempty"`

    // 漏洞描述
    Description   string `json:"description,omitempty"`

}
