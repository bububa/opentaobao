package moziacl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询角色下包含的权限列表 APIResponse
alibaba.mozi.acl.permission.page.rolepermission

根据传入的角色name，分页查询该角色包含的权限列表
*/
type AlibabaMoziAclPermissionPageRolepermissionAPIResponse struct {
    model.CommonResponse
    AlibabaMoziAclPermissionPageRolepermissionResponse
}

type AlibabaMoziAclPermissionPageRolepermissionResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_acl_permission_page_rolepermission_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询角色下权限列表结果对象
    
    Result   *PageRolePermissionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
