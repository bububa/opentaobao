package lstvending

// MultiResultDTO 
type MultiResultDTO struct {
    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 执行成功结果集
    ModuleList   []VendingCargoSpaceDTO `json:"module_list,omitempty" xml:"module_list>vending_cargo_space_dto,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 执行失败结果集
    ErrorList   []AlibabaLstVendingCargospaceSaveResultDTO `json:"error_list,omitempty" xml:"error_list>alibaba_lst_vending_cargospace_save_result_dto,omitempty"`
    // 执行是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
