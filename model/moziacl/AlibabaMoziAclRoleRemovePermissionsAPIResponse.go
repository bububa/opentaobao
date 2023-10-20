package moziacl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclRoleRemovePermissionsAPIResponse 角色移除功能权限 API返回值
// alibaba.mozi.acl.role.remove.permissions
//
// 从角色中移除一批功能权限
type AlibabaMoziAclRoleRemovePermissionsAPIResponse struct {
	model.CommonResponse
	AlibabaMoziAclRoleRemovePermissionsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziAclRoleRemovePermissionsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziAclRoleRemovePermissionsAPIResponseModel).Reset()
}

// AlibabaMoziAclRoleRemovePermissionsAPIResponseModel is 角色移除功能权限 成功返回结果
type AlibabaMoziAclRoleRemovePermissionsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_acl_role_remove_permissions_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 角色移除功能权限结果对象
	Result *RemovePermissionsFromRoleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziAclRoleRemovePermissionsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziAclRoleRemovePermissionsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziAclRoleRemovePermissionsAPIResponse)
	},
}

// GetAlibabaMoziAclRoleRemovePermissionsAPIResponse 从 sync.Pool 获取 AlibabaMoziAclRoleRemovePermissionsAPIResponse
func GetAlibabaMoziAclRoleRemovePermissionsAPIResponse() *AlibabaMoziAclRoleRemovePermissionsAPIResponse {
	return poolAlibabaMoziAclRoleRemovePermissionsAPIResponse.Get().(*AlibabaMoziAclRoleRemovePermissionsAPIResponse)
}

// ReleaseAlibabaMoziAclRoleRemovePermissionsAPIResponse 将 AlibabaMoziAclRoleRemovePermissionsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziAclRoleRemovePermissionsAPIResponse(v *AlibabaMoziAclRoleRemovePermissionsAPIResponse) {
	v.Reset()
	poolAlibabaMoziAclRoleRemovePermissionsAPIResponse.Put(v)
}
