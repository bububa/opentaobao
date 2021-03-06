package kclub

// KcQaFilter 结构体
type KcQaFilter struct {
	// 视角
	Views []int64 `json:"views,omitempty" xml:"views>int64,omitempty"`
	// 状态
	Statuses []int64 `json:"statuses,omitempty" xml:"statuses>int64,omitempty"`
	// 是否需要子知识
	NeedChildKnowledge bool `json:"need_child_knowledge,omitempty" xml:"need_child_knowledge,omitempty"`
}
