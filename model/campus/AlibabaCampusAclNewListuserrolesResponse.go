package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询并标记用户选择的角色 APIResponse
alibaba.campus.acl.new.listuserroles

查询并标记用户选择的角色
*/
type AlibabaCampusAclNewListuserrolesAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewListuserrolesResponse
}

type AlibabaCampusAclNewListuserrolesResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_listuserroles_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
