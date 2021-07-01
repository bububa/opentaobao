package shenjing

// WorkBenchContext 结构体
type WorkBenchContext struct {
	// 授权的appCode
	AppCode string `json:"app_code,omitempty" xml:"app_code,omitempty"`
	// 公司ID
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 系统ID
	SystemId string `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 园区Id
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
}
