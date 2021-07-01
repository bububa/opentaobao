package moziacl

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziAclPermissionpkgAddRolesAPIResponse
将角色添加到权限套餐中 API返回值
alibaba.mozi.acl.permissionpkg.add.roles

此接口是将应用下的一批角色添加到该应用的某个权限套餐中 */
type AlibabaMoziAclPermissionpkgAddRolesAPIResponse struct {
	model.CommonResponse
	AlibabaMoziAclPermissionpkgAddRolesAPIResponseModel
}

// AlibabaMoziAclPermissionpkgAddRolesAPIResponseModel is 将角色添加到权限套餐中 成功返回结果
type AlibabaMoziAclPermissionpkgAddRolesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_acl_permissionpkg_add_roles_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *UpdateRolesToPermissionPackageResult `json:"result,omitempty" xml:"result,omitempty"`
}
