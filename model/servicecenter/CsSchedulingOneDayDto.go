package servicecenter

// CsSchedulingOneDayDto 
/* model for simplify = false
type CsSchedulingOneDayDto struct {

    // 排班日期
    
    Date   string `json:"date,omitempty"`
    

    // 一天排班信息
    
    Schedulings  struct {
        CsSchedulingDto  []CsSchedulingDto `json:"cs_scheduling_dto,omitempty"`
    } `json:"schedulings,omitempty"`
    

}
*/

// CsSchedulingOneDayDto 
type CsSchedulingOneDayDto struct {

    // 排班日期
    Date   string `json:"date,omitempty"`

    // 一天排班信息
    Schedulings   []CsSchedulingDto `json:"schedulings,omitempty"`

}
