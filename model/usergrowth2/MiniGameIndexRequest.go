package usergrowth2

// MiniGameIndexRequest 结构体
type MiniGameIndexRequest struct {
	// 组建集id集合
	CollectionIds []string `json:"collection_ids,omitempty" xml:"collection_ids>string,omitempty"`
	// 活动id
	ActId string `json:"act_id,omitempty" xml:"act_id,omitempty"`
	// 业务场景
	BizScene string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty"`
	// 游戏账号id,仅成语场景使用，该场景下必填
	GameAccId string `json:"game_acc_id,omitempty" xml:"game_acc_id,omitempty"`
	// openId，非成语场景使用，必填
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}
