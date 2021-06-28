package security

// RiskSummary 
/* model for simplify = false
type RiskSummary struct {

    // 仿冒应用信息
    
    FakeInfo  *struct {
        FakeAppSummary  *FakeAppSummary `json:"fake_app_summary,omitempty"`
    } `json:"fake_info,omitempty"`
    

    // 恶意代码信息
    
    MalwareInfo  *struct {
        MalwareSummary  *MalwareSummary `json:"malware_summary,omitempty"`
    } `json:"malware_info,omitempty"`
    

    // 插件信息
    
    PluginInfo  *struct {
        PluginSummary  *PluginSummary `json:"plugin_summary,omitempty"`
    } `json:"plugin_info,omitempty"`
    

    // 任务总状态: 1-已完成,2-处理中,3-处理失败,4-处理超时
    
    TaskStatus   int64 `json:"task_status,omitempty"`
    

    // 漏洞信息
    
    VulnInfo  *struct {
        VulnSummary  *VulnSummary `json:"vuln_summary,omitempty"`
    } `json:"vuln_info,omitempty"`
    

}
*/

// RiskSummary 
type RiskSummary struct {

    // 仿冒应用信息
    FakeInfo   *FakeAppSummary `json:"fake_info,omitempty"`

    // 恶意代码信息
    MalwareInfo   *MalwareSummary `json:"malware_info,omitempty"`

    // 插件信息
    PluginInfo   *PluginSummary `json:"plugin_info,omitempty"`

    // 任务总状态: 1-已完成,2-处理中,3-处理失败,4-处理超时
    TaskStatus   int64 `json:"task_status,omitempty"`

    // 漏洞信息
    VulnInfo   *VulnSummary `json:"vuln_info,omitempty"`

}
