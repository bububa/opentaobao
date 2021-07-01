package damai

// ThirdPerformPushOpenParam 结构体
type ThirdPerformPushOpenParam struct {
	// 场次结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 场次名称
	PerformName string `json:"perform_name,omitempty" xml:"perform_name,omitempty"`
	// 项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 场次开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 场馆id
	VenueId int64 `json:"venue_id,omitempty" xml:"venue_id,omitempty"`
}
