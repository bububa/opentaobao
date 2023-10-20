package drugtrace

import (
	"sync"
)

// RecoverCodeMovePackagingResultDto 结构体
type RecoverCodeMovePackagingResultDto struct {
	// 替换父码
	SourceParentCode string `json:"source_parent_code,omitempty" xml:"source_parent_code,omitempty"`
	// 替换子码
	SourceChildCode string `json:"source_child_code,omitempty" xml:"source_child_code,omitempty"`
	// 被替换父码
	TargetParentCode string `json:"target_parent_code,omitempty" xml:"target_parent_code,omitempty"`
	// 被替换子码
	TargetChildCode string `json:"target_child_code,omitempty" xml:"target_child_code,omitempty"`
	// 码拼箱状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 码拼箱信息
	Info string `json:"info,omitempty" xml:"info,omitempty"`
}

var poolRecoverCodeMovePackagingResultDto = sync.Pool{
	New: func() any {
		return new(RecoverCodeMovePackagingResultDto)
	},
}

// GetRecoverCodeMovePackagingResultDto() 从对象池中获取RecoverCodeMovePackagingResultDto
func GetRecoverCodeMovePackagingResultDto() *RecoverCodeMovePackagingResultDto {
	return poolRecoverCodeMovePackagingResultDto.Get().(*RecoverCodeMovePackagingResultDto)
}

// ReleaseRecoverCodeMovePackagingResultDto 释放RecoverCodeMovePackagingResultDto
func ReleaseRecoverCodeMovePackagingResultDto(v *RecoverCodeMovePackagingResultDto) {
	v.SourceParentCode = ""
	v.SourceChildCode = ""
	v.TargetParentCode = ""
	v.TargetChildCode = ""
	v.Status = ""
	v.Info = ""
	poolRecoverCodeMovePackagingResultDto.Put(v)
}
