package campus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
冻结角色 API返回值 
alibaba.campus.acl.new.freezerole

冻结角色
*/
type AlibabaCampusAclNewFreezeroleAPIResponse struct {
    model.CommonResponse
    AlibabaCampusAclNewFreezeroleAPIResponseModel
}

// 冻结角色 成功返回结果
type AlibabaCampusAclNewFreezeroleAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_campus_acl_new_freezerole_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
