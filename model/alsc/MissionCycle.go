package alsc

// MissionCycle 结构体
type MissionCycle struct {
	// 周期类型
	CycleType string `json:"cycle_type,omitempty" xml:"cycle_type,omitempty"`
	// 任务周期类型扩展（1 天 或固定时间下划线分割开始时间和结束时间 2017-10-01 00:00:00_2017-10-03 00:00:00）
	CycleExt string `json:"cycle_ext,omitempty" xml:"cycle_ext,omitempty"`
	// 完成后是否继续领取 0 对应页面&#34;否&#34;选项，1 对应页面 &#34;是&#34; 选项
	LoopTask int64 `json:"loop_task,omitempty" xml:"loop_task,omitempty"`
}
