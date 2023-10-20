package lstspeacker

import (
	"sync"
)

// AlibabaLstSpeakerFileUploadResultDto 结构体
type AlibabaLstSpeakerFileUploadResultDto struct {
	// 错误码
	ErroMessage string `json:"erro_message,omitempty" xml:"erro_message,omitempty"`
	// 错误码
	ErroCode string `json:"erro_code,omitempty" xml:"erro_code,omitempty"`
	// SpeakerFileDto
	Module *SpeakerFileDto `json:"module,omitempty" xml:"module,omitempty"`
	// 执行结果
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

var poolAlibabaLstSpeakerFileUploadResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstSpeakerFileUploadResultDto)
	},
}

// GetAlibabaLstSpeakerFileUploadResultDto() 从对象池中获取AlibabaLstSpeakerFileUploadResultDto
func GetAlibabaLstSpeakerFileUploadResultDto() *AlibabaLstSpeakerFileUploadResultDto {
	return poolAlibabaLstSpeakerFileUploadResultDto.Get().(*AlibabaLstSpeakerFileUploadResultDto)
}

// ReleaseAlibabaLstSpeakerFileUploadResultDto 释放AlibabaLstSpeakerFileUploadResultDto
func ReleaseAlibabaLstSpeakerFileUploadResultDto(v *AlibabaLstSpeakerFileUploadResultDto) {
	v.ErroMessage = ""
	v.ErroCode = ""
	v.Module = nil
	v.Succ = false
	poolAlibabaLstSpeakerFileUploadResultDto.Put(v)
}
