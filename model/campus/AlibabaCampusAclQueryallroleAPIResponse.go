package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclQueryallroleAPIResponse 查询全部角色 API返回值
// alibaba.campus.acl.queryallrole
//
// 查询全部园区
type AlibabaCampusAclQueryallroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclQueryallroleAPIResponseModel
}

// AlibabaCampusAclQueryallroleAPIResponseModel is 查询全部角色 成功返回结果
type AlibabaCampusAclQueryallroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_queryallrole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CollectionResult `json:"result,omitempty" xml:"result,omitempty"`
}
