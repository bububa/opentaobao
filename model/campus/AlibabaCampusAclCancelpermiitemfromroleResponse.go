package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消角色和权限之间的关系 API返回值 
alibaba.campus.acl.cancelpermiitemfromrole

取消角色和权限之间的关系
*/
type AlibabaCampusAclCancelpermiitemfromroleAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclCancelpermiitemfromroleResponse
}

// 取消角色和权限之间的关系 成功返回结果
type AlibabaCampusAclCancelpermiitemfromroleResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_cancelpermiitemfromrole_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
