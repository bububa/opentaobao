package moziacl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增一个角色 APIResponse
alibaba.mozi.acl.role.add

新增一个角色
*/
type AlibabaMoziAclRoleAddAPIResponse struct {
    model.CommonResponse
    AlibabaMoziAclRoleAddResponse
}

type AlibabaMoziAclRoleAddResponse struct {
    XMLName xml.Name `xml:"alibaba_mozi_acl_role_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 创建角色结果对象
    
    Result   *CreateRoleResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
