package usergrowth2

// PushSendRequest 结构体
type PushSendRequest struct {
	// 活动id
	ActId string `json:"act_id,omitempty" xml:"act_id,omitempty"`
	// 扩展参数
	ExtParams string `json:"ext_params,omitempty" xml:"ext_params,omitempty"`
	// 业务场景
	BizScene string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty"`
	// 给哪个端发push
	SendClient string `json:"send_client,omitempty" xml:"send_client,omitempty"`
	// 游戏账号id,成语场景下必填
	GameAccId string `json:"game_acc_id,omitempty" xml:"game_acc_id,omitempty"`
	// 触达消息模板id
	NotifyId string `json:"notify_id,omitempty" xml:"notify_id,omitempty"`
	// openId，非成语场景下必填
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}
