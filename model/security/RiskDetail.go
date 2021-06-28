package security

// RiskDetail 
/* model for simplify = false
type RiskDetail struct {

    // 任务状态: 1-已完成,2-处理中,3-处理失败,4-处理超时
    
    TaskStatus   int64 `json:"task_status,omitempty"`
    

    // 漏洞信息
    
    VulnInfo  *struct {
        VulnFullInfo  *VulnFullInfo `json:"vuln_full_info,omitempty"`
    } `json:"vuln_info,omitempty"`
    

    // 恶意代码信息
    
    MalwareInfo  *struct {
        MalwareFullInfo  *MalwareFullInfo `json:"malware_full_info,omitempty"`
    } `json:"malware_info,omitempty"`
    

    // 仿冒应用信息
    
    FakeInfo  *struct {
        FakeAppFullInfo  *FakeAppFullInfo `json:"fake_app_full_info,omitempty"`
    } `json:"fake_info,omitempty"`
    

    // 插件信息
    
    PluginInfo  *struct {
        PluginFullInfo  *PluginFullInfo `json:"plugin_full_info,omitempty"`
    } `json:"plugin_info,omitempty"`
    

}
*/

// RiskDetail 
type RiskDetail struct {

    // 任务状态: 1-已完成,2-处理中,3-处理失败,4-处理超时
    TaskStatus   int64 `json:"task_status,omitempty"`

    // 漏洞信息
    VulnInfo   *VulnFullInfo `json:"vuln_info,omitempty"`

    // 恶意代码信息
    MalwareInfo   *MalwareFullInfo `json:"malware_info,omitempty"`

    // 仿冒应用信息
    FakeInfo   *FakeAppFullInfo `json:"fake_info,omitempty"`

    // 插件信息
    PluginInfo   *PluginFullInfo `json:"plugin_info,omitempty"`

}
