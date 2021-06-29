package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
撤销用户授予的角色 APIResponse
alibaba.campus.acl.cancelrolesfromuser

撤销用户授予的角色
*/
type AlibabaCampusAclCancelrolesfromuserAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclCancelrolesfromuserResponse
}

type AlibabaCampusAclCancelrolesfromuserResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_cancelrolesfromuser_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
