package wlbimports

import (
	"sync"
)

// BigbagLogisticsTrajectoryResponse 结构体
type BigbagLogisticsTrajectoryResponse struct {
	// 大包物流轨迹集合
	LogisticsTrajectoryInfoList []BigBagLogisticsTrajectoryInfoResponse `json:"logistics_trajectory_info_list,omitempty" xml:"logistics_trajectory_info_list>big_bag_logistics_trajectory_info_response,omitempty"`
}

var poolBigbagLogisticsTrajectoryResponse = sync.Pool{
	New: func() any {
		return new(BigbagLogisticsTrajectoryResponse)
	},
}

// GetBigbagLogisticsTrajectoryResponse() 从对象池中获取BigbagLogisticsTrajectoryResponse
func GetBigbagLogisticsTrajectoryResponse() *BigbagLogisticsTrajectoryResponse {
	return poolBigbagLogisticsTrajectoryResponse.Get().(*BigbagLogisticsTrajectoryResponse)
}

// ReleaseBigbagLogisticsTrajectoryResponse 释放BigbagLogisticsTrajectoryResponse
func ReleaseBigbagLogisticsTrajectoryResponse(v *BigbagLogisticsTrajectoryResponse) {
	v.LogisticsTrajectoryInfoList = v.LogisticsTrajectoryInfoList[:0]
	poolBigbagLogisticsTrajectoryResponse.Put(v)
}
