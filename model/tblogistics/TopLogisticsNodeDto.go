package tblogistics

// TopLogisticsNodeDto 结构体
type TopLogisticsNodeDto struct {
	// 节点描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 节点枚举
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 当前节点发生时间
	StatusTime int64 `json:"status_time,omitempty" xml:"status_time,omitempty"`
}
