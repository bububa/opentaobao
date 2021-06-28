package kclub

// KcQaFilter 
/* model for simplify = false
type KcQaFilter struct {

    // 是否需要子知识
    
    NeedChildKnowledge   bool `json:"need_child_knowledge,omitempty"`
    

    // 视角
    
    Views  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"views,omitempty"`
    

    // 状态
    
    Statuses  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"statuses,omitempty"`
    

}
*/

// KcQaFilter 
type KcQaFilter struct {

    // 是否需要子知识
    NeedChildKnowledge   bool `json:"need_child_knowledge,omitempty"`

    // 视角
    Views   []int64 `json:"views,omitempty"`

    // 状态
    Statuses   []int64 `json:"statuses,omitempty"`

}
