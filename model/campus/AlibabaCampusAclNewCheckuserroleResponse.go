package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有角色 APIResponse
alibaba.campus.acl.new.checkuserrole

校验用户是否有角色
*/
type AlibabaCampusAclNewCheckuserroleAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewCheckuserroleResponse
}

type AlibabaCampusAclNewCheckuserroleResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_checkuserrole_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回结果
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
