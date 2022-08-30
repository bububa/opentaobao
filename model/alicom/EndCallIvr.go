package alicom

// EndCallIvr 结构体
type EndCallIvr struct {
	// 挂机ivr开关
	EndCallIvr string `json:"end_call_ivr,omitempty" xml:"end_call_ivr,omitempty"`
	// 第一步放音文件
	Step1File string `json:"step1_file,omitempty" xml:"step1_file,omitempty"`
	// 第二步放音文件
	Step2File string `json:"step2_file,omitempty" xml:"step2_file,omitempty"`
	// 有效按键
	ValidKey string `json:"valid_key,omitempty" xml:"valid_key,omitempty"`
	// 最大等待时长，单位秒
	WaitingDtmfTime int64 `json:"waiting_dtmf_time,omitempty" xml:"waiting_dtmf_time,omitempty"`
	// 循环次数
	MaxLoop int64 `json:"max_loop,omitempty" xml:"max_loop,omitempty"`
	// 挂机等待时长
	WaitingEndCall int64 `json:"waiting_end_call,omitempty" xml:"waiting_end_call,omitempty"`
}
