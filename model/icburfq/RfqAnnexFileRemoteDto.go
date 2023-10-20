package icburfq

import (
	"sync"
)

// RfqAnnexFileRemoteDto 结构体
type RfqAnnexFileRemoteDto struct {
	// 文件名
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 唯一文件名
	UniqueFileName string `json:"unique_file_name,omitempty" xml:"unique_file_name,omitempty"`
}

var poolRfqAnnexFileRemoteDto = sync.Pool{
	New: func() any {
		return new(RfqAnnexFileRemoteDto)
	},
}

// GetRfqAnnexFileRemoteDto() 从对象池中获取RfqAnnexFileRemoteDto
func GetRfqAnnexFileRemoteDto() *RfqAnnexFileRemoteDto {
	return poolRfqAnnexFileRemoteDto.Get().(*RfqAnnexFileRemoteDto)
}

// ReleaseRfqAnnexFileRemoteDto 释放RfqAnnexFileRemoteDto
func ReleaseRfqAnnexFileRemoteDto(v *RfqAnnexFileRemoteDto) {
	v.FileName = ""
	v.UniqueFileName = ""
	poolRfqAnnexFileRemoteDto.Put(v)
}
