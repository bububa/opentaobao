package usergrowth2

// PushBatchSendMetaReq 结构体
type PushBatchSendMetaReq struct {
	// 游戏账号id，成语场景下必填
	GameAccId string `json:"game_acc_id,omitempty" xml:"game_acc_id,omitempty"`
	// 模板参数
	TemplateArgs string `json:"template_args,omitempty" xml:"template_args,omitempty"`
	// openId，非成语场景必填
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}
