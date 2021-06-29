package moziacl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
将角色添加到权限套餐中 APIResponse
alibaba.mozi.acl.permissionpkg.add.roles

此接口是将应用下的一批角色添加到该应用的某个权限套餐中
*/
type AlibabaMoziAclPermissionpkgAddRolesAPIResponse struct {
    model.CommonResponse
    AlibabaMoziAclPermissionpkgAddRolesResponse
}

type AlibabaMoziAclPermissionpkgAddRolesResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_acl_permissionpkg_add_roles_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *UpdateRolesToPermissionPackageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
