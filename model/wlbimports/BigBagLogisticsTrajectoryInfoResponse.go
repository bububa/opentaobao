package wlbimports

import (
	"sync"
)

// BigBagLogisticsTrajectoryInfoResponse 结构体
type BigBagLogisticsTrajectoryInfoResponse struct {
	// 操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 操作节点描述（PU-已揽收；OK-已送达；DD-破损送达）
	OperateDesc string `json:"operate_desc,omitempty" xml:"operate_desc,omitempty"`
	// 运单号
	TrackingNumber string `json:"tracking_number,omitempty" xml:"tracking_number,omitempty"`
	// 操作节点（PU、OK、DD）
	OperateCode string `json:"operate_code,omitempty" xml:"operate_code,omitempty"`
}

var poolBigBagLogisticsTrajectoryInfoResponse = sync.Pool{
	New: func() any {
		return new(BigBagLogisticsTrajectoryInfoResponse)
	},
}

// GetBigBagLogisticsTrajectoryInfoResponse() 从对象池中获取BigBagLogisticsTrajectoryInfoResponse
func GetBigBagLogisticsTrajectoryInfoResponse() *BigBagLogisticsTrajectoryInfoResponse {
	return poolBigBagLogisticsTrajectoryInfoResponse.Get().(*BigBagLogisticsTrajectoryInfoResponse)
}

// ReleaseBigBagLogisticsTrajectoryInfoResponse 释放BigBagLogisticsTrajectoryInfoResponse
func ReleaseBigBagLogisticsTrajectoryInfoResponse(v *BigBagLogisticsTrajectoryInfoResponse) {
	v.OperateTime = ""
	v.OperateDesc = ""
	v.TrackingNumber = ""
	v.OperateCode = ""
	poolBigBagLogisticsTrajectoryInfoResponse.Put(v)
}
