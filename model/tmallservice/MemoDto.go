package tmallservice

// MemoDto 结构体
type MemoDto struct {
	// 备注内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 备注创建时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
}
