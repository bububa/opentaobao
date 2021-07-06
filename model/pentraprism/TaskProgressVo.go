package pentraprism

// TaskProgressVo 结构体
type TaskProgressVo struct {
	// 任务完成时间
	FinishedTime string `json:"finished_time,omitempty" xml:"finished_time,omitempty"`
	// 任务状态，“INIT”为初始化状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 任务冷却时间，-1标识无冷却
	CdTime int64 `json:"cd_time,omitempty" xml:"cd_time,omitempty"`
	// 扩展性任务索引
	Index int64 `json:"index,omitempty" xml:"index,omitempty"`
	// 触发多少次任务算完成
	LoopTimes int64 `json:"loop_times,omitempty" xml:"loop_times,omitempty"`
	// 任务完成最大上限
	MaxTimes int64 `json:"max_times,omitempty" xml:"max_times,omitempty"`
	// 还需要做多少次任务才能做完
	NeedTimes int64 `json:"need_times,omitempty" xml:"need_times,omitempty"`
	// 任务周期
	Period int64 `json:"period,omitempty" xml:"period,omitempty"`
	// 任务已经触发的次数
	Times int64 `json:"times,omitempty" xml:"times,omitempty"`
	// 是否达到任务上限
	ReachLimit bool `json:"reach_limit,omitempty" xml:"reach_limit,omitempty"`
}
