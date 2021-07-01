package alihealthcrm

// BatchRemindRequestDto 结构体
type BatchRemindRequestDto struct {
	// 提醒宝宝ID
	BabyId int64 `json:"baby_id,omitempty" xml:"baby_id,omitempty"`
	// 模板参数，可选,|分割
	Params string `json:"params,omitempty" xml:"params,omitempty"`
	// 用户健康ID
	TpUserId int64 `json:"tp_user_id,omitempty" xml:"tp_user_id,omitempty"`
	// 访问跳转链接
	VisitUrl string `json:"visit_url,omitempty" xml:"visit_url,omitempty"`
}
