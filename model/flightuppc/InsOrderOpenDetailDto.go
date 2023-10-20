package flightuppc

import (
	"sync"
)

// InsOrderOpenDetailDto 结构体
type InsOrderOpenDetailDto struct {
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 子保单号
	PolicyNo string `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	// claimApplyTime
	ClaimApplyTime string `json:"claim_apply_time,omitempty" xml:"claim_apply_time,omitempty"`
	// 保险产品编号
	ProductNo string `json:"product_no,omitempty" xml:"product_no,omitempty"`
	// 保司名称
	InsCompany string `json:"ins_company,omitempty" xml:"ins_company,omitempty"`
	// claimSuccessTime
	ClaimSuccessTime string `json:"claim_success_time,omitempty" xml:"claim_success_time,omitempty"`
	// effectiveEndTime
	EffectiveEndTime string `json:"effective_end_time,omitempty" xml:"effective_end_time,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// effectiveStartTime
	EffectiveStartTime string `json:"effective_start_time,omitempty" xml:"effective_start_time,omitempty"`
	// 被保人id
	InsPersonId int64 `json:"ins_person_id,omitempty" xml:"ins_person_id,omitempty"`
	// 保险订单号
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
	// 被保人信息
	InsOrderOpenPerson *InsOrderOpenPersonDto `json:"ins_order_open_person,omitempty" xml:"ins_order_open_person,omitempty"`
	// 保险价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// itemSourceTag
	ItemSourceTag int64 `json:"item_source_tag,omitempty" xml:"item_source_tag,omitempty"`
	// claimFee
	ClaimFee int64 `json:"claim_fee,omitempty" xml:"claim_fee,omitempty"`
	// insSegmentId
	InsSegmentId int64 `json:"ins_segment_id,omitempty" xml:"ins_segment_id,omitempty"`
	// 航段信息
	InsOrderOpenSegment *InsOrderOpenSegmentDto `json:"ins_order_open_segment,omitempty" xml:"ins_order_open_segment,omitempty"`
	// 外部订单号
	OutOrderId int64 `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 子保单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 订单是否有效
	IsEnable bool `json:"is_enable,omitempty" xml:"is_enable,omitempty"`
}

var poolInsOrderOpenDetailDto = sync.Pool{
	New: func() any {
		return new(InsOrderOpenDetailDto)
	},
}

// GetInsOrderOpenDetailDto() 从对象池中获取InsOrderOpenDetailDto
func GetInsOrderOpenDetailDto() *InsOrderOpenDetailDto {
	return poolInsOrderOpenDetailDto.Get().(*InsOrderOpenDetailDto)
}

// ReleaseInsOrderOpenDetailDto 释放InsOrderOpenDetailDto
func ReleaseInsOrderOpenDetailDto(v *InsOrderOpenDetailDto) {
	v.GmtModified = ""
	v.PayTime = ""
	v.PolicyNo = ""
	v.ClaimApplyTime = ""
	v.ProductNo = ""
	v.InsCompany = ""
	v.ClaimSuccessTime = ""
	v.EffectiveEndTime = ""
	v.GmtCreate = ""
	v.EffectiveStartTime = ""
	v.InsPersonId = 0
	v.TcOrderId = 0
	v.InsOrderOpenPerson = nil
	v.Price = 0
	v.ItemSourceTag = 0
	v.ClaimFee = 0
	v.InsSegmentId = 0
	v.InsOrderOpenSegment = nil
	v.OutOrderId = 0
	v.Status = 0
	v.IsEnable = false
	poolInsOrderOpenDetailDto.Put(v)
}
