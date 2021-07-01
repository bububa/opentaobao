package servicecenter

// CsSchedulingOneDayDto 
type CsSchedulingOneDayDto struct {
    // 排班日期
    Date   string `json:"date,omitempty" xml:"date,omitempty"`
    // 一天排班信息
    Schedulings   []CsSchedulingDto `json:"schedulings,omitempty" xml:"schedulings>cs_scheduling_dto,omitempty"`
}
