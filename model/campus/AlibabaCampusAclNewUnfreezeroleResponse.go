package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
解冻角色 APIResponse
alibaba.campus.acl.new.unfreezerole

解冻角色
*/
type AlibabaCampusAclNewUnfreezeroleAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewUnfreezeroleResponse
}

type AlibabaCampusAclNewUnfreezeroleResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_unfreezerole_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
