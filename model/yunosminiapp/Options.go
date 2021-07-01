package yunosminiapp

// Options 结构体
type Options struct {
	// 请求来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 活动步骤
	Step string `json:"step,omitempty" xml:"step,omitempty"`
}
