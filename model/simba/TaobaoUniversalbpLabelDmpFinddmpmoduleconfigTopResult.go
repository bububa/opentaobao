package simba

// TaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult 结构体
type TaobaoUniversalbpLabelDmpFinddmpmoduleconfigTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	DmpModuleConfigVO *DmpModuleConfigVo `json:"dmp_module_config_v_o,omitempty" xml:"dmp_module_config_v_o,omitempty"`
}
