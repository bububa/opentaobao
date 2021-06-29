package moziacl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
角色添加功能权限 APIResponse
alibaba.mozi.acl.role.add.permissions

往角色中添加一批功能权限
*/
type AlibabaMoziAclRoleAddPermissionsAPIResponse struct {
    model.CommonResponse
    AlibabaMoziAclRoleAddPermissionsResponse
}

type AlibabaMoziAclRoleAddPermissionsResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_acl_role_add_permissions_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 角色移除功能权限结果对象
    
    Result   *AddPermissionToRoleResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
