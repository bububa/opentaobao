package drugtrace

import (
	"sync"
)

// CodeMovePackagingResultDto 结构体
type CodeMovePackagingResultDto struct {
	// 替换码
	SourceCode string `json:"source_code,omitempty" xml:"source_code,omitempty"`
	// 被替换码
	TargetCode string `json:"target_code,omitempty" xml:"target_code,omitempty"`
	// 拼箱状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 拼箱信息
	Info string `json:"info,omitempty" xml:"info,omitempty"`
}

var poolCodeMovePackagingResultDto = sync.Pool{
	New: func() any {
		return new(CodeMovePackagingResultDto)
	},
}

// GetCodeMovePackagingResultDto() 从对象池中获取CodeMovePackagingResultDto
func GetCodeMovePackagingResultDto() *CodeMovePackagingResultDto {
	return poolCodeMovePackagingResultDto.Get().(*CodeMovePackagingResultDto)
}

// ReleaseCodeMovePackagingResultDto 释放CodeMovePackagingResultDto
func ReleaseCodeMovePackagingResultDto(v *CodeMovePackagingResultDto) {
	v.SourceCode = ""
	v.TargetCode = ""
	v.Status = ""
	v.Info = ""
	poolCodeMovePackagingResultDto.Put(v)
}
