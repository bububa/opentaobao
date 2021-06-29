package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除管理员 APIResponse
alibaba.campus.acl.new.deleteuserrole

删除管理员
*/
type AlibabaCampusAclNewDeleteuserroleAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewDeleteuserroleResponse
}

type AlibabaCampusAclNewDeleteuserroleResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_deleteuserrole_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
