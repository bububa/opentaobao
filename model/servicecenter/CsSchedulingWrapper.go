package servicecenter

// CsSchedulingWrapper 
type CsSchedulingWrapper struct {

    // 排班记录条数（按天计算）
    TotalCount   int64 `json:"total_count,omitempty"`

    // 按天排班信息
    CsSchedulings   []CsScheduling `json:"cs_schedulings,omitempty"`

}
