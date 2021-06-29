package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
给人直接授权 APIResponse
alibaba.campus.acl.grantpermiitemstouser

给人直接授权
*/
type AlibabaCampusAclGrantpermiitemstouserAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclGrantpermiitemstouserResponse
}

type AlibabaCampusAclGrantpermiitemstouserResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_grantpermiitemstouser_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
