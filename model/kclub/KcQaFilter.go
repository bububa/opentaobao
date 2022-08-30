package kclub

// KcQaFilter 结构体
type KcQaFilter struct {
	// 视角
	Views []string `json:"views,omitempty" xml:"views>string,omitempty"`
	// 状态
	Statuses []string `json:"statuses,omitempty" xml:"statuses>string,omitempty"`
	// 是否需要子知识
	NeedChildKnowledge bool `json:"need_child_knowledge,omitempty" xml:"need_child_knowledge,omitempty"`
}
