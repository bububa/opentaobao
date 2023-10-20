package btrip

import (
	"sync"
)

// BtripExceedApplyRq 结构体
type BtripExceedApplyRq struct {
	// 第三方流程实例id
	ThirdpartyFlowId string `json:"thirdparty_flow_id,omitempty" xml:"thirdparty_flow_id,omitempty"`
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 审批人第三方用户id，多值逗号分隔
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 审批意见
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 商旅超标审批单号
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 审批单状态，1：同意 2：拒绝
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 审批单业务类型，3：超标审批，5：机票改签审批，6：机票退票审批，默认为超标审批
	BizCategory int64 `json:"biz_category,omitempty" xml:"biz_category,omitempty"`
}

var poolBtripExceedApplyRq = sync.Pool{
	New: func() any {
		return new(BtripExceedApplyRq)
	},
}

// GetBtripExceedApplyRq() 从对象池中获取BtripExceedApplyRq
func GetBtripExceedApplyRq() *BtripExceedApplyRq {
	return poolBtripExceedApplyRq.Get().(*BtripExceedApplyRq)
}

// ReleaseBtripExceedApplyRq 释放BtripExceedApplyRq
func ReleaseBtripExceedApplyRq(v *BtripExceedApplyRq) {
	v.ThirdpartyFlowId = ""
	v.CorpId = ""
	v.UserId = ""
	v.Remark = ""
	v.ApplyId = 0
	v.Status = 0
	v.BizCategory = 0
	poolBtripExceedApplyRq.Put(v)
}
