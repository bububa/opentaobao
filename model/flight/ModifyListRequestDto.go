package flight

// ModifyListRequestDto 
/* model for simplify = false
type ModifyListRequestDto struct {

    // 店铺id
    
    AgentIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"agent_ids,omitempty"`
    

    // 申请结束时间
    
    EndApplyTime   string `json:"end_apply_time,omitempty"`
    

    // 申请开始时间
    
    BeginApplyTime   string `json:"begin_apply_time,omitempty"`
    

    // 页码
    
    Page   int64 `json:"page,omitempty"`
    

    // 改签单状态
    
    Status   int64 `json:"status,omitempty"`
    

}
*/

// ModifyListRequestDto 
type ModifyListRequestDto struct {

    // 店铺id
    AgentIds   []int64 `json:"agent_ids,omitempty"`

    // 申请结束时间
    EndApplyTime   string `json:"end_apply_time,omitempty"`

    // 申请开始时间
    BeginApplyTime   string `json:"begin_apply_time,omitempty"`

    // 页码
    Page   int64 `json:"page,omitempty"`

    // 改签单状态
    Status   int64 `json:"status,omitempty"`

}
