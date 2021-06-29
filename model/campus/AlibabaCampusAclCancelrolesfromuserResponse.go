package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
撤销用户授予的角色 API返回值 
alibaba.campus.acl.cancelrolesfromuser

撤销用户授予的角色
*/
type AlibabaCampusAclCancelrolesfromuserAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclCancelrolesfromuserResponse
}

// 撤销用户授予的角色 成功返回结果
type AlibabaCampusAclCancelrolesfromuserResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_cancelrolesfromuser_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
