package drugtrace

import (
	"sync"
)

// SaveCodeRelationResultDto 结构体
type SaveCodeRelationResultDto struct {
	// 结果信息
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 结果标记
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
}

var poolSaveCodeRelationResultDto = sync.Pool{
	New: func() any {
		return new(SaveCodeRelationResultDto)
	},
}

// GetSaveCodeRelationResultDto() 从对象池中获取SaveCodeRelationResultDto
func GetSaveCodeRelationResultDto() *SaveCodeRelationResultDto {
	return poolSaveCodeRelationResultDto.Get().(*SaveCodeRelationResultDto)
}

// ReleaseSaveCodeRelationResultDto 释放SaveCodeRelationResultDto
func ReleaseSaveCodeRelationResultDto(v *SaveCodeRelationResultDto) {
	v.ResponseMessage = ""
	v.Flag = ""
	poolSaveCodeRelationResultDto.Put(v)
}
