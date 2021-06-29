package moziacl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
权限套餐添加权限 APIResponse
alibaba.mozi.acl.permissionpkg.add.permissions

此接口的功能为：将一批应用下的权限添加到该应用下的权限套餐中
*/
type AlibabaMoziAclPermissionpkgAddPermissionsAPIResponse struct {
    model.CommonResponse
    AlibabaMoziAclPermissionpkgAddPermissionsResponse
}

type AlibabaMoziAclPermissionpkgAddPermissionsResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_acl_permissionpkg_add_permissions_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *UpdatePermissionsToPermissionPackageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
