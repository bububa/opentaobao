package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclGrantpermiitemstouserAPIResponse 给人直接授权 API返回值
// alibaba.campus.acl.grantpermiitemstouser
//
// 给人直接授权
type AlibabaCampusAclGrantpermiitemstouserAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclGrantpermiitemstouserAPIResponseModel
}

// AlibabaCampusAclGrantpermiitemstouserAPIResponseModel is 给人直接授权 成功返回结果
type AlibabaCampusAclGrantpermiitemstouserAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_grantpermiitemstouser_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
