package wdk

import (
	"sync"
)

// MedicineResultDto 结构体
type MedicineResultDto struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolMedicineResultDto = sync.Pool{
	New: func() any {
		return new(MedicineResultDto)
	},
}

// GetMedicineResultDto() 从对象池中获取MedicineResultDto
func GetMedicineResultDto() *MedicineResultDto {
	return poolMedicineResultDto.Get().(*MedicineResultDto)
}

// ReleaseMedicineResultDto 释放MedicineResultDto
func ReleaseMedicineResultDto(v *MedicineResultDto) {
	v.ErrorMsg = ""
	v.Success = false
	poolMedicineResultDto.Put(v)
}
