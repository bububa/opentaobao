package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclGetpermissionbyroleidAPIResponse
根据角色Id查询权限 API返回值
alibaba.campus.acl.getpermissionbyroleid

根据角色查询权限 */
type AlibabaCampusAclGetpermissionbyroleidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclGetpermissionbyroleidAPIResponseModel
}

// AlibabaCampusAclGetpermissionbyroleidAPIResponseModel is 根据角色Id查询权限 成功返回结果
type AlibabaCampusAclGetpermissionbyroleidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_getpermissionbyroleid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CollectionResult `json:"result,omitempty" xml:"result,omitempty"`
}
