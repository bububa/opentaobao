package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询全部角色 APIResponse
alibaba.campus.acl.new.listroles

查询全部角色
*/
type AlibabaCampusAclNewListrolesAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewListrolesResponse
}

type AlibabaCampusAclNewListrolesResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_listroles_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
