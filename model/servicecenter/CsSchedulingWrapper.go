package servicecenter

// CsSchedulingWrapper 
/* model for simplify = false
type CsSchedulingWrapper struct {

    // 排班记录条数（按天计算）
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 按天排班信息
    
    CsSchedulings  struct {
        CsScheduling  []CsScheduling `json:"cs_scheduling,omitempty"`
    } `json:"cs_schedulings,omitempty"`
    

}
*/

// CsSchedulingWrapper 
type CsSchedulingWrapper struct {

    // 排班记录条数（按天计算）
    TotalCount   int64 `json:"total_count,omitempty"`

    // 按天排班信息
    CsSchedulings   []CsScheduling `json:"cs_schedulings,omitempty"`

}
