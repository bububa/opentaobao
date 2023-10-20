package moziacl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclPermissionpkgAddPermissionsAPIResponse 权限套餐添加权限 API返回值
// alibaba.mozi.acl.permissionpkg.add.permissions
//
// 此接口的功能为：将一批应用下的权限添加到该应用下的权限套餐中
type AlibabaMoziAclPermissionpkgAddPermissionsAPIResponse struct {
	model.CommonResponse
	AlibabaMoziAclPermissionpkgAddPermissionsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziAclPermissionpkgAddPermissionsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziAclPermissionpkgAddPermissionsAPIResponseModel).Reset()
}

// AlibabaMoziAclPermissionpkgAddPermissionsAPIResponseModel is 权限套餐添加权限 成功返回结果
type AlibabaMoziAclPermissionpkgAddPermissionsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_acl_permissionpkg_add_permissions_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *UpdatePermissionsToPermissionPackageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziAclPermissionpkgAddPermissionsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziAclPermissionpkgAddPermissionsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziAclPermissionpkgAddPermissionsAPIResponse)
	},
}

// GetAlibabaMoziAclPermissionpkgAddPermissionsAPIResponse 从 sync.Pool 获取 AlibabaMoziAclPermissionpkgAddPermissionsAPIResponse
func GetAlibabaMoziAclPermissionpkgAddPermissionsAPIResponse() *AlibabaMoziAclPermissionpkgAddPermissionsAPIResponse {
	return poolAlibabaMoziAclPermissionpkgAddPermissionsAPIResponse.Get().(*AlibabaMoziAclPermissionpkgAddPermissionsAPIResponse)
}

// ReleaseAlibabaMoziAclPermissionpkgAddPermissionsAPIResponse 将 AlibabaMoziAclPermissionpkgAddPermissionsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziAclPermissionpkgAddPermissionsAPIResponse(v *AlibabaMoziAclPermissionpkgAddPermissionsAPIResponse) {
	v.Reset()
	poolAlibabaMoziAclPermissionpkgAddPermissionsAPIResponse.Put(v)
}
