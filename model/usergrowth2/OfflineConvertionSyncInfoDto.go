package usergrowth2

// OfflineConvertionSyncInfoDto 结构体
type OfflineConvertionSyncInfoDto struct {
	// 任务类型
	TaskType string `json:"task_type,omitempty" xml:"task_type,omitempty"`
	// 日期
	Ds int64 `json:"ds,omitempty" xml:"ds,omitempty"`
	// 数据是否更新完毕
	IsSynchronized bool `json:"is_synchronized,omitempty" xml:"is_synchronized,omitempty"`
}
