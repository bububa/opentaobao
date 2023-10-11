package ascp

// UserBlacklistRequest 结构体
type UserBlacklistRequest struct {
	// 黑名单用户
	Blacklist []Blacklist `json:"blacklist,omitempty" xml:"blacklist>blacklist,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 服务类型（1-上门取退）
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 能力（1-上门取退）
	AbilityType string `json:"ability_type,omitempty" xml:"ability_type,omitempty"`
	// 时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 更新类型 1-新增(增量新增) /更新 ；2-删除
	OperateType int64 `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
}
