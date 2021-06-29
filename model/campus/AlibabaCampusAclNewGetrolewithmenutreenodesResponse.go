package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据角色id查询权限 API返回值 
alibaba.campus.acl.new.getrolewithmenutreenodes

根据角色id查询权限
*/
type AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewGetrolewithmenutreenodesResponse
}

// 根据角色id查询权限 成功返回结果
type AlibabaCampusAclNewGetrolewithmenutreenodesResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_getrolewithmenutreenodes_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
