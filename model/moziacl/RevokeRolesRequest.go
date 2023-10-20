package moziacl

import (
	"sync"
)

// RevokeRolesRequest 结构体
type RevokeRolesRequest struct {
	// 回收的角色name列表
	RoleNames []string `json:"role_names,omitempty" xml:"role_names>string,omitempty"`
	// 回收的角色所在的应用的app name
	TargetAppName string `json:"target_app_name,omitempty" xml:"target_app_name,omitempty"`
	// 请求扩展字段
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 回收主体对象
	Principal *BucUserPrincipalParam `json:"principal,omitempty" xml:"principal,omitempty"`
}

var poolRevokeRolesRequest = sync.Pool{
	New: func() any {
		return new(RevokeRolesRequest)
	},
}

// GetRevokeRolesRequest() 从对象池中获取RevokeRolesRequest
func GetRevokeRolesRequest() *RevokeRolesRequest {
	return poolRevokeRolesRequest.Get().(*RevokeRolesRequest)
}

// ReleaseRevokeRolesRequest 释放RevokeRolesRequest
func ReleaseRevokeRolesRequest(v *RevokeRolesRequest) {
	v.RoleNames = v.RoleNames[:0]
	v.TargetAppName = ""
	v.RequestMetaData = ""
	v.Principal = nil
	poolRevokeRolesRequest.Put(v)
}
