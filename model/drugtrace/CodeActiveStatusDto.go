package drugtrace

import (
	"sync"
)

// CodeActiveStatusDto 结构体
type CodeActiveStatusDto struct {
	// 码段
	ResProdCode string `json:"res_prod_code,omitempty" xml:"res_prod_code,omitempty"`
	// 最后激活时间，到毫秒时间timeMills方式
	ActiveTime int64 `json:"active_time,omitempty" xml:"active_time,omitempty"`
}

var poolCodeActiveStatusDto = sync.Pool{
	New: func() any {
		return new(CodeActiveStatusDto)
	},
}

// GetCodeActiveStatusDto() 从对象池中获取CodeActiveStatusDto
func GetCodeActiveStatusDto() *CodeActiveStatusDto {
	return poolCodeActiveStatusDto.Get().(*CodeActiveStatusDto)
}

// ReleaseCodeActiveStatusDto 释放CodeActiveStatusDto
func ReleaseCodeActiveStatusDto(v *CodeActiveStatusDto) {
	v.ResProdCode = ""
	v.ActiveTime = 0
	poolCodeActiveStatusDto.Put(v)
}
