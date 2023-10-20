package lstvending

// MultiResultDto 结构体
type MultiResultDto struct {
	// 执行成功结果集
	ModuleList []VendingCargoSpaceDto `json:"module_list,omitempty" xml:"module_list>vending_cargo_space_dto,omitempty"`
	// 执行失败结果集
	ErrorList []AlibabalstvendingcargospacesaveResultDto `json:"error_list,omitempty" xml:"error_list>alibabalstvendingcargospacesave_result_dto,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
