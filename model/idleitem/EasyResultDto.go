package idleitem

import (
	"sync"
)

// EasyResultDto 结构体
type EasyResultDto struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 上传成功的文件id
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功与否
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolEasyResultDto = sync.Pool{
	New: func() any {
		return new(EasyResultDto)
	},
}

// GetEasyResultDto() 从对象池中获取EasyResultDto
func GetEasyResultDto() *EasyResultDto {
	return poolEasyResultDto.Get().(*EasyResultDto)
}

// ReleaseEasyResultDto 释放EasyResultDto
func ReleaseEasyResultDto(v *EasyResultDto) {
	v.ErrorCode = ""
	v.Data = ""
	v.ErrorMsg = ""
	v.Success = false
	poolEasyResultDto.Put(v)
}
