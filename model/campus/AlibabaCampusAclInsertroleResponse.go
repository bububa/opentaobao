package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增角色 APIResponse
alibaba.campus.acl.insertrole

新增角色
*/
type AlibabaCampusAclInsertroleAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclInsertroleResponse
}

type AlibabaCampusAclInsertroleResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_insertrole_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *RoleRsp `json:"result,omitempty" xml:"result,omitempty"`

    
}
