package alsc

// MiniGameGrantIntegralRequest 结构体
type MiniGameGrantIntegralRequest struct {
	// 组建集id集合
	CollectionIds []string `json:"collection_ids,omitempty" xml:"collection_ids>string,omitempty"`
	// 活动id
	ActId string `json:"act_id,omitempty" xml:"act_id,omitempty"`
	// 业务场景
	BizScene string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty"`
	// 游戏账号，成语场景下必填
	GameAccId string `json:"game_acc_id,omitempty" xml:"game_acc_id,omitempty"`
	// 请求id(幂等id),不超过128位,相同幂等号,发放数量也要保持相同，否则会幂等异常
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 资产id
	PropertyId string `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// openId，非成语场景下必填
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 积分发放数量
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}
