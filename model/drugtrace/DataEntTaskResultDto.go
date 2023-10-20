package drugtrace

import (
	"sync"
)

// DataEntTaskResultDto 结构体
type DataEntTaskResultDto struct {
	// 诊断原因
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDataEntTaskResultDto = sync.Pool{
	New: func() any {
		return new(DataEntTaskResultDto)
	},
}

// GetDataEntTaskResultDto() 从对象池中获取DataEntTaskResultDto
func GetDataEntTaskResultDto() *DataEntTaskResultDto {
	return poolDataEntTaskResultDto.Get().(*DataEntTaskResultDto)
}

// ReleaseDataEntTaskResultDto 释放DataEntTaskResultDto
func ReleaseDataEntTaskResultDto(v *DataEntTaskResultDto) {
	v.Model = ""
	v.MsgInfo = ""
	v.HttpStatusCode = 0
	v.Success = false
	poolDataEntTaskResultDto.Put(v)
}
