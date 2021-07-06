package jstinteractive

// InteractiveTask 结构体
type InteractiveTask struct {
	// 任务ID
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 任务状态，INIT=初始化完成，待领取；ACCEPTED=已领取 正在做；AWARDING=领奖中；CURRENT_COMPLETE=本次任务完成；COMPLETE=已完成
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 任务列表素材
	Material *Material `json:"material,omitempty" xml:"material,omitempty"`
	// 任务列表进度
	Progress *Progress `json:"progress,omitempty" xml:"progress,omitempty"`
	// 任务类型，1=商品详情页，2=直播间
	TaskType int64 `json:"task_type,omitempty" xml:"task_type,omitempty"`
	// 当前任务是否在进行中，默认为true。直播任务中的false代表当前不在直播时间段，不返回action字段，前端应屏蔽跳转
	InProgress bool `json:"in_progress,omitempty" xml:"in_progress,omitempty"`
}
