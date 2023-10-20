package tmallgenie

import (
	"sync"
)

// DeviceCorpusTopDto 结构体
type DeviceCorpusTopDto struct {
	// 操作语料
	CorpusList []string `json:"corpus_list,omitempty" xml:"corpus_list>string,omitempty"`
	// 支持的操作类型
	FunctionName string `json:"function_name,omitempty" xml:"function_name,omitempty"`
}

var poolDeviceCorpusTopDto = sync.Pool{
	New: func() any {
		return new(DeviceCorpusTopDto)
	},
}

// GetDeviceCorpusTopDto() 从对象池中获取DeviceCorpusTopDto
func GetDeviceCorpusTopDto() *DeviceCorpusTopDto {
	return poolDeviceCorpusTopDto.Get().(*DeviceCorpusTopDto)
}

// ReleaseDeviceCorpusTopDto 释放DeviceCorpusTopDto
func ReleaseDeviceCorpusTopDto(v *DeviceCorpusTopDto) {
	v.CorpusList = v.CorpusList[:0]
	v.FunctionName = ""
	poolDeviceCorpusTopDto.Put(v)
}
