package rail

// RailCarrierRs 
type RailCarrierRs struct {
    // 是否成功
    Success   string `json:"success,omitempty" xml:"success,omitempty"`
    // 错误描述
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 错误码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 铁路运营公司列表
    ModuleList   []RailCarrierRS `json:"module_list,omitempty" xml:"module_list>rail_carrier_rs,omitempty"`
}
