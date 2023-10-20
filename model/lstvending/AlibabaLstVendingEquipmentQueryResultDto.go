package lstvending

// AlibabalstvendingequipmentqueryResultDto 结构体
type AlibabalstvendingequipmentqueryResultDto struct {
	// 设备信息列表
	ModuleList []OpenEquipmentDto `json:"module_list,omitempty" xml:"module_list>open_equipment_dto,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否执行异常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
