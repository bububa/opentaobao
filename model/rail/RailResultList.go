package rail

// RailResultList 
type RailResultList struct {
    // 是否成功
    Success   string `json:"success,omitempty" xml:"success,omitempty"`
    // 错误描述
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 错误码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 城市列表
    ModuleList   []RailDivisionRs `json:"module_list,omitempty" xml:"module_list>rail_division_rs,omitempty"`
}
