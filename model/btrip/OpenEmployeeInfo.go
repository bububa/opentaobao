package btrip

import (
	"sync"
)

// OpenEmployeeInfo 结构体
type OpenEmployeeInfo struct {
	// 员工昵称。
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 第三方员工ID。
	ThirdPartUserId string `json:"third_part_user_id,omitempty" xml:"third_part_user_id,omitempty"`
	// 第三方员工工号。
	ThirdPartJobNo string `json:"third_part_job_no,omitempty" xml:"third_part_job_no,omitempty"`
	// 1:离职 0:在职
	LeaveStatus int64 `json:"leave_status,omitempty" xml:"leave_status,omitempty"`
}

var poolOpenEmployeeInfo = sync.Pool{
	New: func() any {
		return new(OpenEmployeeInfo)
	},
}

// GetOpenEmployeeInfo() 从对象池中获取OpenEmployeeInfo
func GetOpenEmployeeInfo() *OpenEmployeeInfo {
	return poolOpenEmployeeInfo.Get().(*OpenEmployeeInfo)
}

// ReleaseOpenEmployeeInfo 释放OpenEmployeeInfo
func ReleaseOpenEmployeeInfo(v *OpenEmployeeInfo) {
	v.UserName = ""
	v.ThirdPartUserId = ""
	v.ThirdPartJobNo = ""
	v.LeaveStatus = 0
	poolOpenEmployeeInfo.Put(v)
}
