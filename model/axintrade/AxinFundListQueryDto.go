package axintrade

import (
	"sync"
)

// AxinFundListQueryDto 结构体
type AxinFundListQueryDto struct {
	// 资金单类型
	FundStatus []string `json:"fund_status,omitempty" xml:"fund_status>string,omitempty"`
	// 外部订单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
}

var poolAxinFundListQueryDto = sync.Pool{
	New: func() any {
		return new(AxinFundListQueryDto)
	},
}

// GetAxinFundListQueryDto() 从对象池中获取AxinFundListQueryDto
func GetAxinFundListQueryDto() *AxinFundListQueryDto {
	return poolAxinFundListQueryDto.Get().(*AxinFundListQueryDto)
}

// ReleaseAxinFundListQueryDto 释放AxinFundListQueryDto
func ReleaseAxinFundListQueryDto(v *AxinFundListQueryDto) {
	v.FundStatus = v.FundStatus[:0]
	v.OuterOrderId = ""
	poolAxinFundListQueryDto.Put(v)
}
