package campus

import (
	"sync"
)

// WorkBenchContext 结构体
type WorkBenchContext struct {
	// 应用id
	SystemId string `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// app编码
	AppCode string `json:"app_code,omitempty" xml:"app_code,omitempty"`
	// securityCode
	SecurityCode string `json:"security_code,omitempty" xml:"security_code,omitempty"`
	// campusCode
	CampusCode string `json:"campus_code,omitempty" xml:"campus_code,omitempty"`
	// userName
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// eagleEyeTraceId
	EagleEyeTraceId string `json:"eagle_eye_trace_id,omitempty" xml:"eagle_eye_trace_id,omitempty"`
	// ip
	Ip string `json:"ip,omitempty" xml:"ip,omitempty"`
	// 登录者的语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 公司id
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 园区id
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// userId
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolWorkBenchContext = sync.Pool{
	New: func() any {
		return new(WorkBenchContext)
	},
}

// GetWorkBenchContext() 从对象池中获取WorkBenchContext
func GetWorkBenchContext() *WorkBenchContext {
	return poolWorkBenchContext.Get().(*WorkBenchContext)
}

// ReleaseWorkBenchContext 释放WorkBenchContext
func ReleaseWorkBenchContext(v *WorkBenchContext) {
	v.SystemId = ""
	v.AppCode = ""
	v.SecurityCode = ""
	v.CampusCode = ""
	v.UserName = ""
	v.EagleEyeTraceId = ""
	v.Ip = ""
	v.Language = ""
	v.CompanyId = 0
	v.CampusId = 0
	v.UserId = 0
	poolWorkBenchContext.Put(v)
}
