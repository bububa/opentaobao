package moziacl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclRoleAddPermissionsAPIResponse 角色添加功能权限 API返回值
// alibaba.mozi.acl.role.add.permissions
//
// 往角色中添加一批功能权限
type AlibabaMoziAclRoleAddPermissionsAPIResponse struct {
	model.CommonResponse
	AlibabaMoziAclRoleAddPermissionsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziAclRoleAddPermissionsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziAclRoleAddPermissionsAPIResponseModel).Reset()
}

// AlibabaMoziAclRoleAddPermissionsAPIResponseModel is 角色添加功能权限 成功返回结果
type AlibabaMoziAclRoleAddPermissionsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_acl_role_add_permissions_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 角色移除功能权限结果对象
	Result *AddPermissionToRoleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziAclRoleAddPermissionsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziAclRoleAddPermissionsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziAclRoleAddPermissionsAPIResponse)
	},
}

// GetAlibabaMoziAclRoleAddPermissionsAPIResponse 从 sync.Pool 获取 AlibabaMoziAclRoleAddPermissionsAPIResponse
func GetAlibabaMoziAclRoleAddPermissionsAPIResponse() *AlibabaMoziAclRoleAddPermissionsAPIResponse {
	return poolAlibabaMoziAclRoleAddPermissionsAPIResponse.Get().(*AlibabaMoziAclRoleAddPermissionsAPIResponse)
}

// ReleaseAlibabaMoziAclRoleAddPermissionsAPIResponse 将 AlibabaMoziAclRoleAddPermissionsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziAclRoleAddPermissionsAPIResponse(v *AlibabaMoziAclRoleAddPermissionsAPIResponse) {
	v.Reset()
	poolAlibabaMoziAclRoleAddPermissionsAPIResponse.Put(v)
}
