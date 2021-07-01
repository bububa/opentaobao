package pentraprism

// TaskRewardVo 结构体
type TaskRewardVo struct {
	// 查询奖励错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 查询奖励详细错误信息
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 发奖励时间
	Time int64 `json:"time,omitempty" xml:"time,omitempty"`
	// 发奖励类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 是否领奖
	Win bool `json:"win,omitempty" xml:"win,omitempty"`
	// 奖励详细信息
	Result *TaskRewardItemVo `json:"result,omitempty" xml:"result,omitempty"`
}
