package alsc

// Stage 结构体
type Stage struct {
	// 阶段金额 阶段金额值 	count 和sum必须要有一个有值
	Sum string `json:"sum,omitempty" xml:"sum,omitempty"`
	// 阶段数 阶段金额值 	count 和sum必须要有一个有值
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
