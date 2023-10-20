package alihouse

import (
	"sync"
)

// ProjectSalesTimePrePermitDto 结构体
type ProjectSalesTimePrePermitDto struct {
	// 预售证外部id
	PreSalePermitId string `json:"pre_sale_permit_id,omitempty" xml:"pre_sale_permit_id,omitempty"`
	// 预售时间
	PreSalePermitTime string `json:"pre_sale_permit_time,omitempty" xml:"pre_sale_permit_time,omitempty"`
	// 预售证编号
	PreSalePermitNo string `json:"pre_sale_permit_no,omitempty" xml:"pre_sale_permit_no,omitempty"`
}

var poolProjectSalesTimePrePermitDto = sync.Pool{
	New: func() any {
		return new(ProjectSalesTimePrePermitDto)
	},
}

// GetProjectSalesTimePrePermitDto() 从对象池中获取ProjectSalesTimePrePermitDto
func GetProjectSalesTimePrePermitDto() *ProjectSalesTimePrePermitDto {
	return poolProjectSalesTimePrePermitDto.Get().(*ProjectSalesTimePrePermitDto)
}

// ReleaseProjectSalesTimePrePermitDto 释放ProjectSalesTimePrePermitDto
func ReleaseProjectSalesTimePrePermitDto(v *ProjectSalesTimePrePermitDto) {
	v.PreSalePermitId = ""
	v.PreSalePermitTime = ""
	v.PreSalePermitNo = ""
	poolProjectSalesTimePrePermitDto.Put(v)
}
