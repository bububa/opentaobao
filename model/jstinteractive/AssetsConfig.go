package jstinteractive

// AssetsConfig 结构体
type AssetsConfig struct {
	// 做任务按钮，默认值【去完成】
	AcceptBtn string `json:"accept_btn,omitempty" xml:"accept_btn,omitempty"`
	// 待领奖按钮，默认值【领取奖励】
	AwardBtn string `json:"award_btn,omitempty" xml:"award_btn,omitempty"`
	// 任务完成按钮，默认值【已完成】
	CompleteBtn string `json:"complete_btn,omitempty" xml:"complete_btn,omitempty"`
	// 任务描述，默认值为title
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 任务图标，图片大小80x80
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 领取任务按钮，默认值【领取任务】
	InitBtn string `json:"init_btn,omitempty" xml:"init_btn,omitempty"`
	// 任务副标题，默认值为title
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 任务id
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 任务标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品详情页地址或直播间地址
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 任务浮标标题，浏览商品为【浏览10秒得奖励】，观看直播为【观看30秒得奖励】，暂不支持自定义
	ShopCompTitle string `json:"shop_comp_title,omitempty" xml:"shop_comp_title,omitempty"`
	// 任务浏览时长，商品详情页10秒，直播间30秒，暂不支持自定义
	Duration string `json:"duration,omitempty" xml:"duration,omitempty"`
	// 商品id，必须为授权店铺的商品，配置商品浏览任务时必填
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 直播间id，必须为授权店铺的直播间，配置直播观看任务时必填
	VideoId int64 `json:"video_id,omitempty" xml:"video_id,omitempty"`
	// 任务类型，1=浏览商品，1=观看直播
	TaskType int64 `json:"task_type,omitempty" xml:"task_type,omitempty"`
}
