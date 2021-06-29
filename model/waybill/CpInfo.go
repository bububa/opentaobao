package waybill

// CpInfo 
type CpInfo struct {
    // 地址信息
    Address   *Address `json:"address,omitempty" xml:"address,omitempty"`
    // 云打印模板
    CloudTemplateId   string `json:"cloud_template_id,omitempty" xml:"cloud_template_id,omitempty"`
    // 快递公司
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    // 状态: 0-禁用, 1-启用
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
}
