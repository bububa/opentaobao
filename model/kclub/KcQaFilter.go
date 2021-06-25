package kclub

// KcQaFilter 
type KcQaFilter struct {

    // 是否需要子知识
    NeedChildKnowledge   bool `json:"need_child_knowledge,omitempty"`

    // 视角
    Views   []Number `json:"views,omitempty"`

    // 状态
    Statuses   []Number `json:"statuses,omitempty"`

}
