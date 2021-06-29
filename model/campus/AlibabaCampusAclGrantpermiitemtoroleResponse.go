package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
权限赋予角色 APIResponse
alibaba.campus.acl.grantpermiitemtorole

权限赋予角色
*/
type AlibabaCampusAclGrantpermiitemtoroleAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclGrantpermiitemtoroleResponse
}

type AlibabaCampusAclGrantpermiitemtoroleResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_grantpermiitemtorole_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
