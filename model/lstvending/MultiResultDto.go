package lstvending

import (
	"sync"
)

// MultiResultDto 结构体
type MultiResultDto struct {
	// 执行成功结果集
	ModuleList []VendingCargoSpaceDto `json:"module_list,omitempty" xml:"module_list>vending_cargo_space_dto,omitempty"`
	// 执行失败结果集
	ErrorList []AlibabaLstVendingCargospaceSaveResultDto `json:"error_list,omitempty" xml:"error_list>alibaba_lst_vending_cargospace_save_result_dto,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 执行是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolMultiResultDto = sync.Pool{
	New: func() any {
		return new(MultiResultDto)
	},
}

// GetMultiResultDto() 从对象池中获取MultiResultDto
func GetMultiResultDto() *MultiResultDto {
	return poolMultiResultDto.Get().(*MultiResultDto)
}

// ReleaseMultiResultDto 释放MultiResultDto
func ReleaseMultiResultDto(v *MultiResultDto) {
	v.ModuleList = v.ModuleList[:0]
	v.ErrorList = v.ErrorList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Success = false
	poolMultiResultDto.Put(v)
}
