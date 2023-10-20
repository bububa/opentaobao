package alsc

import (
	"sync"
)

// QueryRechargeAccountFlowOpenReq 结构体
type QueryRechargeAccountFlowOpenReq struct {
	// 品牌ID(不能和outbrandid同时为空)
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 流水Id
	FlowId string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
	// 外部品牌ID(不能和brandid同时为空)
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 外部订单ID，和流水ID必传一
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
}

var poolQueryRechargeAccountFlowOpenReq = sync.Pool{
	New: func() any {
		return new(QueryRechargeAccountFlowOpenReq)
	},
}

// GetQueryRechargeAccountFlowOpenReq() 从对象池中获取QueryRechargeAccountFlowOpenReq
func GetQueryRechargeAccountFlowOpenReq() *QueryRechargeAccountFlowOpenReq {
	return poolQueryRechargeAccountFlowOpenReq.Get().(*QueryRechargeAccountFlowOpenReq)
}

// ReleaseQueryRechargeAccountFlowOpenReq 释放QueryRechargeAccountFlowOpenReq
func ReleaseQueryRechargeAccountFlowOpenReq(v *QueryRechargeAccountFlowOpenReq) {
	v.BrandId = ""
	v.FlowId = ""
	v.OutBrandId = ""
	v.OuterOrderId = ""
	poolQueryRechargeAccountFlowOpenReq.Put(v)
}
