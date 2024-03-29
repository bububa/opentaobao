package drug

import (
	"sync"
)

// LogisticsOrderDto 结构体
type LogisticsOrderDto struct {
	// 快递公司编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 快递单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 电子面单
	Waybill string `json:"waybill,omitempty" xml:"waybill,omitempty"`
}

var poolLogisticsOrderDto = sync.Pool{
	New: func() any {
		return new(LogisticsOrderDto)
	},
}

// GetLogisticsOrderDto() 从对象池中获取LogisticsOrderDto
func GetLogisticsOrderDto() *LogisticsOrderDto {
	return poolLogisticsOrderDto.Get().(*LogisticsOrderDto)
}

// ReleaseLogisticsOrderDto 释放LogisticsOrderDto
func ReleaseLogisticsOrderDto(v *LogisticsOrderDto) {
	v.CpCode = ""
	v.MailNo = ""
	v.Waybill = ""
	poolLogisticsOrderDto.Put(v)
}
