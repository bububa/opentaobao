package security

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
