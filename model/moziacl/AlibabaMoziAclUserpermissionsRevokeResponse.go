package moziacl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
回收账户权限 APIResponse
alibaba.mozi.acl.userpermissions.revoke

调用此接口，会根据入参回收该账户下的该批权限点
*/
type AlibabaMoziAclUserpermissionsRevokeAPIResponse struct {
    model.CommonResponse
    AlibabaMoziAclUserpermissionsRevokeResponse
}

type AlibabaMoziAclUserpermissionsRevokeResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_acl_userpermissions_revoke_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 回收账户的权限请求返回对象
    
    Result   *RevokePermissionsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
