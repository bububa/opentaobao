package moziacl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增一个角色 API返回值 
alibaba.mozi.acl.role.add

新增一个角色
*/
type AlibabaMoziAclRoleAddAPIResponse struct {
    model.CommonResponse
    AlibabaMoziAclRoleAddAPIResponseModel
}

// 新增一个角色 成功返回结果
type AlibabaMoziAclRoleAddAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mozi_acl_role_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 创建角色结果对象
    Result   *CreateRoleResult `json:"result,omitempty" xml:"result,omitempty"`
}
