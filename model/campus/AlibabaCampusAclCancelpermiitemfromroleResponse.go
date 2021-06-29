package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消角色和权限之间的关系 APIResponse
alibaba.campus.acl.cancelpermiitemfromrole

取消角色和权限之间的关系
*/
type AlibabaCampusAclCancelpermiitemfromroleAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclCancelpermiitemfromroleResponse
}

type AlibabaCampusAclCancelpermiitemfromroleResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_cancelpermiitemfromrole_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
