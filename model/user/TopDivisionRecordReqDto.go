package user

// TopDivisionRecordReqDto 
type TopDivisionRecordReqDto struct {
    // 页码
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // 每页数量
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 开始时间(毫秒为单位)
    StartTime   int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 结束时间(毫秒为单位)
    EndTime   int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
}
