package logistic

import (
	"sync"
)

// LogisticsResourceDto 结构体
type LogisticsResourceDto struct {
	// 运单号校验正则表达式
	RegMailNo string `json:"reg_mail_no,omitempty" xml:"reg_mail_no,omitempty"`
	// 快递资源编码
	ResourceCode string `json:"resource_code,omitempty" xml:"resource_code,omitempty"`
	// 快递资源名称
	ResourceName string `json:"resource_name,omitempty" xml:"resource_name,omitempty"`
	// 快递公司id
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
}

var poolLogisticsResourceDto = sync.Pool{
	New: func() any {
		return new(LogisticsResourceDto)
	},
}

// GetLogisticsResourceDto() 从对象池中获取LogisticsResourceDto
func GetLogisticsResourceDto() *LogisticsResourceDto {
	return poolLogisticsResourceDto.Get().(*LogisticsResourceDto)
}

// ReleaseLogisticsResourceDto 释放LogisticsResourceDto
func ReleaseLogisticsResourceDto(v *LogisticsResourceDto) {
	v.RegMailNo = ""
	v.ResourceCode = ""
	v.ResourceName = ""
	v.CompanyId = 0
	poolLogisticsResourceDto.Put(v)
}
