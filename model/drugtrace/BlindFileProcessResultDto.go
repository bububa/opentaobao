package drugtrace

import (
	"sync"
)

// BlindFileProcessResultDto 结构体
type BlindFileProcessResultDto struct {
	// 企业名称
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 文件名称
	BlindFileName string `json:"blind_file_name,omitempty" xml:"blind_file_name,omitempty"`
	// 处理状态 (0:处理成功,1:处理失败,2:正在处理,3:待处理)
	ProcessStatus string `json:"process_status,omitempty" xml:"process_status,omitempty"`
	// 失败原因
	FailReason string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
}

var poolBlindFileProcessResultDto = sync.Pool{
	New: func() any {
		return new(BlindFileProcessResultDto)
	},
}

// GetBlindFileProcessResultDto() 从对象池中获取BlindFileProcessResultDto
func GetBlindFileProcessResultDto() *BlindFileProcessResultDto {
	return poolBlindFileProcessResultDto.Get().(*BlindFileProcessResultDto)
}

// ReleaseBlindFileProcessResultDto 释放BlindFileProcessResultDto
func ReleaseBlindFileProcessResultDto(v *BlindFileProcessResultDto) {
	v.RefEntId = ""
	v.BlindFileName = ""
	v.ProcessStatus = ""
	v.FailReason = ""
	poolBlindFileProcessResultDto.Put(v)
}
