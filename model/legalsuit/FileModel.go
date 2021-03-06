package legalsuit

// FileModel 结构体
type FileModel struct {
	// 附件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 附件Key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 沟通时间
	SubmitDate string `json:"submit_date,omitempty" xml:"submit_date,omitempty"`
	// 沟通结果
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 沟通内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 沟通方
	Intro string `json:"intro,omitempty" xml:"intro,omitempty"`
}
