package alsc

import (
	"sync"
)

// PointFlowOpenInfo 结构体
type PointFlowOpenInfo struct {
	// 业务场景
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 业务类型描述
	BizTypeDescription string `json:"biz_type_description,omitempty" xml:"biz_type_description,omitempty"`
	// 流水id
	FlowId string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
	// 操作员名
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 交易单号
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 变更原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 交易时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 变更积分
	ChangePoint int64 `json:"change_point,omitempty" xml:"change_point,omitempty"`
	// 剩余积分
	RemainPoint int64 `json:"remain_point,omitempty" xml:"remain_point,omitempty"`
}

var poolPointFlowOpenInfo = sync.Pool{
	New: func() any {
		return new(PointFlowOpenInfo)
	},
}

// GetPointFlowOpenInfo() 从对象池中获取PointFlowOpenInfo
func GetPointFlowOpenInfo() *PointFlowOpenInfo {
	return poolPointFlowOpenInfo.Get().(*PointFlowOpenInfo)
}

// ReleasePointFlowOpenInfo 释放PointFlowOpenInfo
func ReleasePointFlowOpenInfo(v *PointFlowOpenInfo) {
	v.BizType = ""
	v.BizTypeDescription = ""
	v.FlowId = ""
	v.OperatorName = ""
	v.OutBizId = ""
	v.Reason = ""
	v.ShopName = ""
	v.Time = ""
	v.ChangePoint = 0
	v.RemainPoint = 0
	poolPointFlowOpenInfo.Put(v)
}
