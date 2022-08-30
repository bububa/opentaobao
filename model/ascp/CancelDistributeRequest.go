package ascp

// CancelDistributeRequest 结构体
type CancelDistributeRequest struct {
	// 取消详情
	CancelDistributeInfoList []CancelDistributeInfo `json:"cancel_distribute_info_list,omitempty" xml:"cancel_distribute_info_list>cancel_distribute_info,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间。时间戳。 毫秒
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
