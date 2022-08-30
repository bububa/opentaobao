package cloudgame

// GetSlotResponse 结构体
type GetSlotResponse struct {
	// 联机码
	JoinCode []string `json:"join_code,omitempty" xml:"join_code>string,omitempty"`
	// 串流信息
	Endpoints []GameEndPoint `json:"endpoints,omitempty" xml:"endpoints>game_end_point,omitempty"`
	// h5servier地址
	EndpointServer []EndpointServerDto `json:"endpoint_server,omitempty" xml:"endpoint_server>endpoint_server_dto,omitempty"`
	// 编码后的游戏ID
	MixGameId string `json:"mix_game_id,omitempty" xml:"mix_game_id,omitempty"`
	//  会话ID
	GameSession string `json:"game_session,omitempty" xml:"game_session,omitempty"`
	// 游戏token
	GameToken string `json:"game_token,omitempty" xml:"game_token,omitempty"`
	// 游戏状态描述
	QueueStateDesc string `json:"queue_state_desc,omitempty" xml:"queue_state_desc,omitempty"`
	// 会话状态备注
	LinkMark string `json:"link_mark,omitempty" xml:"link_mark,omitempty"`
	// 用户id
	MixUserId string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
	// 调度类型
	SchedType string `json:"sched_type,omitempty" xml:"sched_type,omitempty"`
	// 会话状态
	QueueState int64 `json:"queue_state,omitempty" xml:"queue_state,omitempty"`
	// 用户等级
	UserLevel int64 `json:"user_level,omitempty" xml:"user_level,omitempty"`
	// vip等级
	VipLevel int64 `json:"vip_level,omitempty" xml:"vip_level,omitempty"`
	// 轮询时间间隔
	PollingInterval int64 `json:"polling_interval,omitempty" xml:"polling_interval,omitempty"`
	// 是否是主控
	Host bool `json:"host,omitempty" xml:"host,omitempty"`
	// 是否是强制264
	ForceH264 bool `json:"force_h264,omitempty" xml:"force_h264,omitempty"`
}
