package tmallnr

// NrZqsPauseInfoDto 
type NrZqsPauseInfoDto struct {
    // 暂停开始时间，包含该天
    PauseStartDay   string `json:"pause_start_day,omitempty" xml:"pause_start_day,omitempty"`
    // 暂停结束时间，包含该天
    PauseEndDay   string `json:"pause_end_day,omitempty" xml:"pause_end_day,omitempty"`
}
