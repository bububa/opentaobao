package jstinteractive

// Material 结构体
type Material struct {
	// 待完成按钮文案
	AcceptBtn string `json:"accept_btn,omitempty" xml:"accept_btn,omitempty"`
	// 任务列表动作类型，VISIT=浏览
	ActionType string `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 任务列表副标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 任务列表图标链接
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 任务列表动作，比如浏览的店铺地址链接
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 任务列表标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 待领取按钮文案
	InitBtn string `json:"init_btn,omitempty" xml:"init_btn,omitempty"`
	// 待领取按钮文案
	AwardBtn string `json:"award_btn,omitempty" xml:"award_btn,omitempty"`
	// 已完成按钮文案
	CompleteBtn string `json:"complete_btn,omitempty" xml:"complete_btn,omitempty"`
	// 浏览任务需要多少秒才能完成
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
}
