package security

// RiskDetailBatchItem 
/* model for simplify = false
type RiskDetailBatchItem struct {

    // app的md5
    
    AppIdentity   string `json:"app_identity,omitempty"`
    

    // 恶意代码检测结果
    
    MalwareInfo  *struct {
        MalwareFullInfo  *MalwareFullInfo `json:"malware_full_info,omitempty"`
    } `json:"malware_info,omitempty"`
    

    // 漏洞列表(任务完成时才返回)
    
    VulnInfo  *struct {
        VulnFullInfo  *VulnFullInfo `json:"vuln_full_info,omitempty"`
    } `json:"vuln_info,omitempty"`
    

    // 插件信息
    
    PluginInfo  *struct {
        PluginFullInfo  *PluginFullInfo `json:"plugin_full_info,omitempty"`
    } `json:"plugin_info,omitempty"`
    

}
*/

// RiskDetailBatchItem 
type RiskDetailBatchItem struct {

    // app的md5
    AppIdentity   string `json:"app_identity,omitempty"`

    // 恶意代码检测结果
    MalwareInfo   *MalwareFullInfo `json:"malware_info,omitempty"`

    // 漏洞列表(任务完成时才返回)
    VulnInfo   *VulnFullInfo `json:"vuln_info,omitempty"`

    // 插件信息
    PluginInfo   *PluginFullInfo `json:"plugin_info,omitempty"`

}
