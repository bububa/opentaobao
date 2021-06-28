package kclub

// KcQaFilter 
type KcQaFilter struct {

    // 是否需要子知识
    
    NeedChildKnowledge   bool `json:"need_child_knowledge,omitempty" xml:"need_child_knowledge,omitempty"`
    

    // 视角
    
    Views   []int64 `json:"views,omitempty" xml:"views>int64,omitempty"`
    

    // 状态
    
    Statuses   []int64 `json:"statuses,omitempty" xml:"statuses>int64,omitempty"`
    

}
