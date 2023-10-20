package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclInsertroleAPIResponse 新增角色 API返回值
// alibaba.campus.acl.insertrole
//
// 新增角色
type AlibabaCampusAclInsertroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclInsertroleAPIResponseModel
}

// AlibabaCampusAclInsertroleAPIResponseModel is 新增角色 成功返回结果
type AlibabaCampusAclInsertroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_insertrole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RoleRsp `json:"result,omitempty" xml:"result,omitempty"`
}
