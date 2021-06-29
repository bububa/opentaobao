package tmallgenie

// ScheduleDto 
type ScheduleDto struct {

    // 一次性
    
    Once   *OnceSchedule `json:"once,omitempty" xml:"once,omitempty"`
    

    // 法定工作日
    
    StatutoryWorkingDay   *StatutoryWorkingDaySchedule `json:"statutory_working_day,omitempty" xml:"statutory_working_day,omitempty"`
    

    // 调度类型
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // 每周
    
    Weekly   *WeeklySchedule `json:"weekly,omitempty" xml:"weekly,omitempty"`
    

}
