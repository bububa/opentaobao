package alihealthoutflow

// AsyncPrescribeSearchRequest 结构体
type AsyncPrescribeSearchRequest struct {
	// 开方结束时间
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 开方开始时间
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 医院id
	HospitalId string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
	// 每页多少条
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 第几页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
}
