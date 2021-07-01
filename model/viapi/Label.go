package viapi

// Label 结构体
type Label struct {
	// 可选值包括：  spam：含垃圾信息 politics： 涉政 abuse：辱骂 porn：智能鉴黄 terrorism：暴恐识别 flood：灌水 contraband：违禁 ad：文本违规识别
	Label string `json:"label,omitempty" xml:"label,omitempty"`
}
