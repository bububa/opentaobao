package moziacl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
角色移除功能权限 APIResponse
alibaba.mozi.acl.role.remove.permissions

从角色中移除一批功能权限
*/
type AlibabaMoziAclRoleRemovePermissionsAPIResponse struct {
    model.CommonResponse
    AlibabaMoziAclRoleRemovePermissionsResponse
}

type AlibabaMoziAclRoleRemovePermissionsResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_acl_role_remove_permissions_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 角色移除功能权限结果对象
    
    Result   *RemovePermissionsFromRoleResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
