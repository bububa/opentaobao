package drugtrace

import (
	"sync"
)

// BillUpstreamDto 结构体
type BillUpstreamDto struct {
	// 发货企业名称
	FromUserName string `json:"from_user_name,omitempty" xml:"from_user_name,omitempty"`
	// 单据时间
	BillTime string `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
	// 货主
	RefUserId string `json:"ref_user_id,omitempty" xml:"ref_user_id,omitempty"`
	// 发货企业ID
	FromUserId string `json:"from_user_id,omitempty" xml:"from_user_id,omitempty"`
	// 单据类型
	BillType string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
	// 收货企业名称
	ToUserName string `json:"to_user_name,omitempty" xml:"to_user_name,omitempty"`
	// 单据号
	BillCode string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
	// 收货企业ID
	ToUserId string `json:"to_user_id,omitempty" xml:"to_user_id,omitempty"`
	// 收货企业REF_ENT_ID
	ToRefUserId string `json:"to_ref_user_id,omitempty" xml:"to_ref_user_id,omitempty"`
	// 发货企业REF_ENT_ID
	FromRefUserId string `json:"from_ref_user_id,omitempty" xml:"from_ref_user_id,omitempty"`
}

var poolBillUpstreamDto = sync.Pool{
	New: func() any {
		return new(BillUpstreamDto)
	},
}

// GetBillUpstreamDto() 从对象池中获取BillUpstreamDto
func GetBillUpstreamDto() *BillUpstreamDto {
	return poolBillUpstreamDto.Get().(*BillUpstreamDto)
}

// ReleaseBillUpstreamDto 释放BillUpstreamDto
func ReleaseBillUpstreamDto(v *BillUpstreamDto) {
	v.FromUserName = ""
	v.BillTime = ""
	v.RefUserId = ""
	v.FromUserId = ""
	v.BillType = ""
	v.ToUserName = ""
	v.BillCode = ""
	v.ToUserId = ""
	v.ToRefUserId = ""
	v.FromRefUserId = ""
	poolBillUpstreamDto.Put(v)
}
