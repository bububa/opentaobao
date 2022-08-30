package alihouse

// TradeToolBindParamDto 结构体
type TradeToolBindParamDto struct {
	// 绑定对象
	OuterBindParamDto []OuterBindParamDto `json:"outer_bind_param_dto,omitempty" xml:"outer_bind_param_dto>outer_bind_param_dto,omitempty"`
	// 绑定类型：1.楼盘，8.房源，6.楼栋，7.户型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 关系是否有效，1-有效，0-无效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 请求时间戳，精确到毫秒
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}
