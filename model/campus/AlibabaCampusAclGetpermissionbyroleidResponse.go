package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据角色Id查询权限 APIResponse
alibaba.campus.acl.getpermissionbyroleid

根据角色查询权限
*/
type AlibabaCampusAclGetpermissionbyroleidAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclGetpermissionbyroleidResponse
}

type AlibabaCampusAclGetpermissionbyroleidResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_getpermissionbyroleid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *CollectionResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
