package campus

import (
	"sync"
)

// RoleRsp 结构体
type RoleRsp struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误级别
	ErrorLevel string `json:"error_level,omitempty" xml:"error_level,omitempty"`
	// 角色id
	RoleId string `json:"role_id,omitempty" xml:"role_id,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolRoleRsp = sync.Pool{
	New: func() any {
		return new(RoleRsp)
	},
}

// GetRoleRsp() 从对象池中获取RoleRsp
func GetRoleRsp() *RoleRsp {
	return poolRoleRsp.Get().(*RoleRsp)
}

// ReleaseRoleRsp 释放RoleRsp
func ReleaseRoleRsp(v *RoleRsp) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.ErrorLevel = ""
	v.RoleId = ""
	v.Success = false
	poolRoleRsp.Put(v)
}
