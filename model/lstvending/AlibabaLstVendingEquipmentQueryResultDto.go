package lstvending

import (
	"sync"
)

// AlibabaLstVendingEquipmentQueryResultDto 结构体
type AlibabaLstVendingEquipmentQueryResultDto struct {
	// 设备信息列表
	ModuleList []OpenEquipmentDto `json:"module_list,omitempty" xml:"module_list>open_equipment_dto,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否执行异常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLstVendingEquipmentQueryResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendingEquipmentQueryResultDto)
	},
}

// GetAlibabaLstVendingEquipmentQueryResultDto() 从对象池中获取AlibabaLstVendingEquipmentQueryResultDto
func GetAlibabaLstVendingEquipmentQueryResultDto() *AlibabaLstVendingEquipmentQueryResultDto {
	return poolAlibabaLstVendingEquipmentQueryResultDto.Get().(*AlibabaLstVendingEquipmentQueryResultDto)
}

// ReleaseAlibabaLstVendingEquipmentQueryResultDto 释放AlibabaLstVendingEquipmentQueryResultDto
func ReleaseAlibabaLstVendingEquipmentQueryResultDto(v *AlibabaLstVendingEquipmentQueryResultDto) {
	v.ModuleList = v.ModuleList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolAlibabaLstVendingEquipmentQueryResultDto.Put(v)
}
