package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有该角色 API返回值 
alibaba.campus.acl.checkemprole

校验用户是否有该权限
*/
type AlibabaCampusAclCheckemproleAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclCheckemproleResponse
}

// 校验用户是否有该角色 成功返回结果
type AlibabaCampusAclCheckemproleResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_checkemprole_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
