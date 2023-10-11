package usergrowth

// TaskConfig 结构体
type TaskConfig struct {
	// 任务等级，可空
	Level string `json:"level,omitempty" xml:"level,omitempty"`
	// 通过安全校验，可空
	PassSafeCheck bool `json:"pass_safe_check,omitempty" xml:"pass_safe_check,omitempty"`
}
