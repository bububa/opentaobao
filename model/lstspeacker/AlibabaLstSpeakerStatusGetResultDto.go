package lstspeacker

import (
	"sync"
)

// AlibabaLstSpeakerStatusGetResultDto 结构体
type AlibabaLstSpeakerStatusGetResultDto struct {
	// 错误码
	ErroMessage string `json:"erro_message,omitempty" xml:"erro_message,omitempty"`
	// 错误码
	ErroCode string `json:"erro_code,omitempty" xml:"erro_code,omitempty"`
	// 状态对象
	Module *SpeakerOnLineStatus `json:"module,omitempty" xml:"module,omitempty"`
	// 执行结果
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

var poolAlibabaLstSpeakerStatusGetResultDto = sync.Pool{
	New: func() any {
		return new(AlibabaLstSpeakerStatusGetResultDto)
	},
}

// GetAlibabaLstSpeakerStatusGetResultDto() 从对象池中获取AlibabaLstSpeakerStatusGetResultDto
func GetAlibabaLstSpeakerStatusGetResultDto() *AlibabaLstSpeakerStatusGetResultDto {
	return poolAlibabaLstSpeakerStatusGetResultDto.Get().(*AlibabaLstSpeakerStatusGetResultDto)
}

// ReleaseAlibabaLstSpeakerStatusGetResultDto 释放AlibabaLstSpeakerStatusGetResultDto
func ReleaseAlibabaLstSpeakerStatusGetResultDto(v *AlibabaLstSpeakerStatusGetResultDto) {
	v.ErroMessage = ""
	v.ErroCode = ""
	v.Module = nil
	v.Succ = false
	poolAlibabaLstSpeakerStatusGetResultDto.Put(v)
}
