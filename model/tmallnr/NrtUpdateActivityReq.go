package tmallnr

// NrtUpdateActivityReq 
type NrtUpdateActivityReq struct {
    // 活动ID
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    // 页面ID
    PageId   int64 `json:"page_id,omitempty" xml:"page_id,omitempty"`
    // 员工ID
    EmployeeId   int64 `json:"employee_id,omitempty" xml:"employee_id,omitempty"`
}
