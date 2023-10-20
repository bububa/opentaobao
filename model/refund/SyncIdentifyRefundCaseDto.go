package refund

import (
	"sync"
)

// SyncIdentifyRefundCaseDto 结构体
type SyncIdentifyRefundCaseDto struct {
	// 鉴定工单操作备注
	OperateTips string `json:"operate_tips,omitempty" xml:"operate_tips,omitempty"`
	// 鉴定工单ID
	OuterCaseId string `json:"outer_case_id,omitempty" xml:"outer_case_id,omitempty"`
	// 扩展属性，json格式
	ExtAttrs string `json:"ext_attrs,omitempty" xml:"ext_attrs,omitempty"`
	// 子订单ID
	DetailOrderId int64 `json:"detail_order_id,omitempty" xml:"detail_order_id,omitempty"`
	// 数据发生时间绝对秒数，如鉴定工单创建时间、鉴定工单完结时间
	OccurTime int64 `json:"occur_time,omitempty" xml:"occur_time,omitempty"`
	// 工单操作类型，1：开启，2：完结
	OperateType int64 `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
	// 退款ID
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
}

var poolSyncIdentifyRefundCaseDto = sync.Pool{
	New: func() any {
		return new(SyncIdentifyRefundCaseDto)
	},
}

// GetSyncIdentifyRefundCaseDto() 从对象池中获取SyncIdentifyRefundCaseDto
func GetSyncIdentifyRefundCaseDto() *SyncIdentifyRefundCaseDto {
	return poolSyncIdentifyRefundCaseDto.Get().(*SyncIdentifyRefundCaseDto)
}

// ReleaseSyncIdentifyRefundCaseDto 释放SyncIdentifyRefundCaseDto
func ReleaseSyncIdentifyRefundCaseDto(v *SyncIdentifyRefundCaseDto) {
	v.OperateTips = ""
	v.OuterCaseId = ""
	v.ExtAttrs = ""
	v.DetailOrderId = 0
	v.OccurTime = 0
	v.OperateType = 0
	v.RefundId = 0
	poolSyncIdentifyRefundCaseDto.Put(v)
}
