package moziacl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
回收账户被授予的角色接口 API返回值 
alibaba.mozi.acl.userroles.revoke

调用此接口，会根据入参回收该账户下的该批角色
*/
type AlibabaMoziAclUserrolesRevokeAPIResponse struct {
    model.CommonResponse
    AlibabaMoziAclUserrolesRevokeResponse
}

// 回收账户被授予的角色接口 成功返回结果
type AlibabaMoziAclUserrolesRevokeResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_acl_userroles_revoke_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 回收角色结果对象
    Result   *RevokeRolesResult `json:"result,omitempty" xml:"result,omitempty"`
}
