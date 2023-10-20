package alihouse

import (
	"sync"
)

// SubAccountReqDto 结构体
type SubAccountReqDto struct {
	// 主账号昵称
	MainUserNick string `json:"main_user_nick,omitempty" xml:"main_user_nick,omitempty"`
	// 昵称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 外部id 由角色类型决定id类型
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 默认密码
	Identity string `json:"identity,omitempty" xml:"identity,omitempty"`
	// 电话号码
	Phone int64 `json:"phone,omitempty" xml:"phone,omitempty"`
	// 是否测试 0-非测试 1-测试
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// etc版本号
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
	// 子账号角色类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 入离职状态 0-未知 1-在职 2-离职
	JobStatus int64 `json:"job_status,omitempty" xml:"job_status,omitempty"`
}

var poolSubAccountReqDto = sync.Pool{
	New: func() any {
		return new(SubAccountReqDto)
	},
}

// GetSubAccountReqDto() 从对象池中获取SubAccountReqDto
func GetSubAccountReqDto() *SubAccountReqDto {
	return poolSubAccountReqDto.Get().(*SubAccountReqDto)
}

// ReleaseSubAccountReqDto 释放SubAccountReqDto
func ReleaseSubAccountReqDto(v *SubAccountReqDto) {
	v.MainUserNick = ""
	v.UserName = ""
	v.OuterConsultantId = ""
	v.Identity = ""
	v.Phone = 0
	v.IsTest = 0
	v.EtcVersion = 0
	v.Type = 0
	v.JobStatus = 0
	poolSubAccountReqDto.Put(v)
}
