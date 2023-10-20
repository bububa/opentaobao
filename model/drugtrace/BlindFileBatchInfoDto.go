package drugtrace

import (
	"sync"
)

// BlindFileBatchInfoDto 结构体
type BlindFileBatchInfoDto struct {
	// 批次号
	BatchNo string `json:"batch_no,omitempty" xml:"batch_no,omitempty"`
	// 生产日期（yyyy-MM-dd HH:mm:ss）
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 有效期至（yyyy-MM-dd HH:mm:ss）
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

var poolBlindFileBatchInfoDto = sync.Pool{
	New: func() any {
		return new(BlindFileBatchInfoDto)
	},
}

// GetBlindFileBatchInfoDto() 从对象池中获取BlindFileBatchInfoDto
func GetBlindFileBatchInfoDto() *BlindFileBatchInfoDto {
	return poolBlindFileBatchInfoDto.Get().(*BlindFileBatchInfoDto)
}

// ReleaseBlindFileBatchInfoDto 释放BlindFileBatchInfoDto
func ReleaseBlindFileBatchInfoDto(v *BlindFileBatchInfoDto) {
	v.BatchNo = ""
	v.ProduceDate = ""
	v.ExpireDate = ""
	poolBlindFileBatchInfoDto.Put(v)
}
