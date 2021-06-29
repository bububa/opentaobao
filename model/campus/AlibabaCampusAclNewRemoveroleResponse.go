package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除角色 API返回值 
alibaba.campus.acl.new.removerole

删除角色
*/
type AlibabaCampusAclNewRemoveroleAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewRemoveroleResponse
}

// 删除角色 成功返回结果
type AlibabaCampusAclNewRemoveroleResponse struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_removerole_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // {}
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
