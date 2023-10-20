package drugtrace

import (
	"sync"
)

// DataEntTaskDto 结构体
type DataEntTaskDto struct {
	// fileInfoList
	FileInfoList []string `json:"file_info_list,omitempty" xml:"file_info_list>string,omitempty"`
	// fileNum
	FileNum int64 `json:"file_num,omitempty" xml:"file_num,omitempty"`
}

var poolDataEntTaskDto = sync.Pool{
	New: func() any {
		return new(DataEntTaskDto)
	},
}

// GetDataEntTaskDto() 从对象池中获取DataEntTaskDto
func GetDataEntTaskDto() *DataEntTaskDto {
	return poolDataEntTaskDto.Get().(*DataEntTaskDto)
}

// ReleaseDataEntTaskDto 释放DataEntTaskDto
func ReleaseDataEntTaskDto(v *DataEntTaskDto) {
	v.FileInfoList = v.FileInfoList[:0]
	v.FileNum = 0
	poolDataEntTaskDto.Put(v)
}
