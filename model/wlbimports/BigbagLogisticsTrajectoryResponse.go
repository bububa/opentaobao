package wlbimports

// BigbagLogisticsTrajectoryResponse 结构体
type BigbagLogisticsTrajectoryResponse struct {
	// 大包物流轨迹集合
	LogisticsTrajectoryInfoList []BigBagLogisticsTrajectoryInfoResponse `json:"logistics_trajectory_info_list,omitempty" xml:"logistics_trajectory_info_list>big_bag_logistics_trajectory_info_response,omitempty"`
}
