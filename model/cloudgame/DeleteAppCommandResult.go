package cloudgame

// DeleteAppCommandResult 结构体
type DeleteAppCommandResult struct {
	// 游戏删除任务ID
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 游戏删除是否成功
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}
