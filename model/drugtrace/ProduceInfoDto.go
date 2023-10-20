package drugtrace

import (
	"sync"
)

// ProduceInfoDto 结构体
type ProduceInfoDto struct {
	// 批次号
	BatchNo string `json:"batch_no,omitempty" xml:"batch_no,omitempty"`
	// 生产日期
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 有效期至
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 最小包装数量
	PkgAmount string `json:"pkg_amount,omitempty" xml:"pkg_amount,omitempty"`
	// 生产日期
	ProduceDateStr string `json:"produce_date_str,omitempty" xml:"produce_date_str,omitempty"`
	// 码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 生产批号
	ProduceBatchNo string `json:"produce_batch_no,omitempty" xml:"produce_batch_no,omitempty"`
}

var poolProduceInfoDto = sync.Pool{
	New: func() any {
		return new(ProduceInfoDto)
	},
}

// GetProduceInfoDto() 从对象池中获取ProduceInfoDto
func GetProduceInfoDto() *ProduceInfoDto {
	return poolProduceInfoDto.Get().(*ProduceInfoDto)
}

// ReleaseProduceInfoDto 释放ProduceInfoDto
func ReleaseProduceInfoDto(v *ProduceInfoDto) {
	v.BatchNo = ""
	v.ProduceDate = ""
	v.ExpireDate = ""
	v.PkgAmount = ""
	v.ProduceDateStr = ""
	v.Code = ""
	v.ProduceBatchNo = ""
	poolProduceInfoDto.Put(v)
}
