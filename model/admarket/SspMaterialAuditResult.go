package admarket

// SspMaterialAuditResult 
type SspMaterialAuditResult struct {
    // 原因
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    // 创意id
    MaterialId   int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
    // 审核结果
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 排除设备
    ExcludeDevices   []ExcludeDevice `json:"exclude_devices,omitempty" xml:"exclude_devices>exclude_device,omitempty"`
    // 渠道
    Channel   string `json:"channel,omitempty" xml:"channel,omitempty"`
}
