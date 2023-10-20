package ascp

import (
	"sync"
)

// PublicFuturePlanRequest 结构体
type PublicFuturePlanRequest struct {
	// 负卖品详情
	FuturePlanDetailList []FuturePlanDetail `json:"future_plan_detail_list,omitempty" xml:"future_plan_detail_list>future_plan_detail,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作id。所有的逻辑请保持品+仓+发布单号唯一，如果相同的品使用了相同的发布单号，库存认为是同一个计划。这个发布单号可以是在途单据号（CO、PO、ICP（调拨单号）、计划单号）
	OperationOrderId string `json:"operation_order_id,omitempty" xml:"operation_order_id,omitempty"`
	// 操作code。幂等逻辑：品+operationOrderId+operationDetailOrderId+operationCode。 如果存在多次操作（期货数量更新） 需要保持operationCode不一样
	OperationCode string `json:"operation_code,omitempty" xml:"operation_code,omitempty"`
	// 新增负卖计划用 FU010,更新用 FU020,停用用 FU030
	BizActivityCode string `json:"biz_activity_code,omitempty" xml:"biz_activity_code,omitempty"`
	// 业务请求时间。时间戳。 毫秒
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolPublicFuturePlanRequest = sync.Pool{
	New: func() any {
		return new(PublicFuturePlanRequest)
	},
}

// GetPublicFuturePlanRequest() 从对象池中获取PublicFuturePlanRequest
func GetPublicFuturePlanRequest() *PublicFuturePlanRequest {
	return poolPublicFuturePlanRequest.Get().(*PublicFuturePlanRequest)
}

// ReleasePublicFuturePlanRequest 释放PublicFuturePlanRequest
func ReleasePublicFuturePlanRequest(v *PublicFuturePlanRequest) {
	v.FuturePlanDetailList = v.FuturePlanDetailList[:0]
	v.RequestId = ""
	v.OperationOrderId = ""
	v.OperationCode = ""
	v.BizActivityCode = ""
	v.RequestTime = 0
	poolPublicFuturePlanRequest.Put(v)
}
