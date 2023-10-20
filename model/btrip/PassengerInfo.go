package btrip

import (
	"sync"
)

// PassengerInfo 结构体
type PassengerInfo struct {
	// 项目code
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 成本中心名称
	CostCenterName string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// 成本中心编号
	CostCenterNumber string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	// 第三方项目id
	ThirdpartProjectId string `json:"thirdpart_project_id,omitempty" xml:"thirdpart_project_id,omitempty"`
	// 用户名
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 项目名称
	ProjectTitle string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// 成本中心id
	CostCenterId int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// 出行人类型 0:内部/1:外部
	UserType int64 `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
}

var poolPassengerInfo = sync.Pool{
	New: func() any {
		return new(PassengerInfo)
	},
}

// GetPassengerInfo() 从对象池中获取PassengerInfo
func GetPassengerInfo() *PassengerInfo {
	return poolPassengerInfo.Get().(*PassengerInfo)
}

// ReleasePassengerInfo 释放PassengerInfo
func ReleasePassengerInfo(v *PassengerInfo) {
	v.ProjectCode = ""
	v.CostCenterName = ""
	v.CostCenterNumber = ""
	v.ThirdpartProjectId = ""
	v.UserName = ""
	v.UserId = ""
	v.ProjectTitle = ""
	v.CostCenterId = 0
	v.UserType = 0
	v.ProjectId = 0
	poolPassengerInfo.Put(v)
}
