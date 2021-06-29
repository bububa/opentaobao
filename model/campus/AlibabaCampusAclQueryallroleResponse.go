package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询全部角色 APIResponse
alibaba.campus.acl.queryallrole

查询全部园区
*/
type AlibabaCampusAclQueryallroleAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclQueryallroleResponse
}

type AlibabaCampusAclQueryallroleResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_queryallrole_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *CollectionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
