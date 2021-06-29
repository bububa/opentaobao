package servicecenter

// CsSchedulingOneDayDTO 
type CsSchedulingOneDayDTO struct {
    // 排班日期
    Date   string `json:"date,omitempty" xml:"date,omitempty"`
    // 一天排班信息
    Schedulings   []CsSchedulingDTO `json:"schedulings,omitempty" xml:"schedulings>cs_scheduling_dto,omitempty"`
}
