package alihouse

import (
	"sync"
)

// ProjectVerifyTicketDto 结构体
type ProjectVerifyTicketDto struct {
	// 补贴赞助方全称
	SubsidySponsor string `json:"subsidy_sponsor,omitempty" xml:"subsidy_sponsor,omitempty"`
	// 返利优惠信息
	RebateRemark string `json:"rebate_remark,omitempty" xml:"rebate_remark,omitempty"`
	// 外部楼盘id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 交易商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 交易商品类型:1-房款立减  2-购房返现 3-其他权益 4-一元一平（这期忽略） 5-百亿房补 6-特价房
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 返利金额（单位:分）
	RebateTotalAmount int64 `json:"rebate_total_amount,omitempty" xml:"rebate_total_amount,omitempty"`
	// 优惠开始时间（单位:毫秒）
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 优惠结束时间（单位:毫秒）
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 购买金额（单位:分）
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 排序权重
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
}

var poolProjectVerifyTicketDto = sync.Pool{
	New: func() any {
		return new(ProjectVerifyTicketDto)
	},
}

// GetProjectVerifyTicketDto() 从对象池中获取ProjectVerifyTicketDto
func GetProjectVerifyTicketDto() *ProjectVerifyTicketDto {
	return poolProjectVerifyTicketDto.Get().(*ProjectVerifyTicketDto)
}

// ReleaseProjectVerifyTicketDto 释放ProjectVerifyTicketDto
func ReleaseProjectVerifyTicketDto(v *ProjectVerifyTicketDto) {
	v.SubsidySponsor = ""
	v.RebateRemark = ""
	v.OuterId = ""
	v.ItemId = 0
	v.Type = 0
	v.RebateTotalAmount = 0
	v.StartTime = 0
	v.EndTime = 0
	v.Amount = 0
	v.Weight = 0
	poolProjectVerifyTicketDto.Put(v)
}
