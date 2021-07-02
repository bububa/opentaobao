package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclGetmenubyempidAPIResponse 查询用户的菜单 API返回值
// alibaba.campus.acl.getmenubyempid
//
// 查询用户的菜单
type AlibabaCampusAclGetmenubyempidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclGetmenubyempidAPIResponseModel
}

// AlibabaCampusAclGetmenubyempidAPIResponseModel is 查询用户的菜单 成功返回结果
type AlibabaCampusAclGetmenubyempidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_getmenubyempid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *CollectionResult `json:"result,omitempty" xml:"result,omitempty"`
}
