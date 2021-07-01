package alihealthmedical

// OuterMsgContent 结构体
type OuterMsgContent struct {
	// 文本
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 音频
	Radio string `json:"radio,omitempty" xml:"radio,omitempty"`
	// 音频时长
	RadioTime int64 `json:"radio_time,omitempty" xml:"radio_time,omitempty"`
	// 诊断
	Diagnose string `json:"diagnose,omitempty" xml:"diagnose,omitempty"`
	// 建议
	Advice string `json:"advice,omitempty" xml:"advice,omitempty"`
	// 发送时间戳
	SendTime int64 `json:"send_time,omitempty" xml:"send_time,omitempty"`
	// 图片
	Pic []string `json:"pic,omitempty" xml:"pic>string,omitempty"`
}
