package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclGrantpermiitemtoroleAPIResponse 权限赋予角色 API返回值
// alibaba.campus.acl.grantpermiitemtorole
//
// 权限赋予角色
type AlibabaCampusAclGrantpermiitemtoroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclGrantpermiitemtoroleAPIResponseModel
}

// AlibabaCampusAclGrantpermiitemtoroleAPIResponseModel is 权限赋予角色 成功返回结果
type AlibabaCampusAclGrantpermiitemtoroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_grantpermiitemtorole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
