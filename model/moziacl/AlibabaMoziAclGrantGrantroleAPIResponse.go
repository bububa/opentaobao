package moziacl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
将一个角色授予一个账号 API返回值 
alibaba.mozi.acl.grant.grantrole

根据入参，将入参中的角色授权给入参的某个账户，调用此接口后，该账户就会被授予该角色
*/
type AlibabaMoziAclGrantGrantroleAPIResponse struct {
    model.CommonResponse
    AlibabaMoziAclGrantGrantroleAPIResponseModel
}

// 将一个角色授予一个账号 成功返回结果
type AlibabaMoziAclGrantGrantroleAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mozi_acl_grant_grantrole_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 给账号授予角色结果
    Result   *GrantRolesResult `json:"result,omitempty" xml:"result,omitempty"`
}
