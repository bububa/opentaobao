package alilabs

// AlibabaAilabsTmallgenieAuthSwitchuserResult 
type AlibabaAilabsTmallgenieAuthSwitchuserResult struct {
    // 扩展信息
    Extension   string `json:"extension,omitempty" xml:"extension,omitempty"`
    // 注册结果
    RegisterInfoVO   *RegisterInfoVo `json:"register_info_v_o,omitempty" xml:"register_info_v_o,omitempty"`
    // 授权结果
    DeviceTokenVO   *DeviceTokenVo `json:"device_token_v_o,omitempty" xml:"device_token_v_o,omitempty"`
}
