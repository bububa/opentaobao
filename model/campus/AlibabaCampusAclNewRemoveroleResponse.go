package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除角色 APIResponse
alibaba.campus.acl.new.removerole

删除角色
*/
type AlibabaCampusAclNewRemoveroleAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewRemoveroleResponse
}

type AlibabaCampusAclNewRemoveroleResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_removerole_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // {}
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
