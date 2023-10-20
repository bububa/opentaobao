package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformserviceabilitydeleteAPIResponse 删除服务能力 API返回值
// alibaba.ssc.supplyplatform.serviceability.delete
//
// 删除服务能力
type AlibabasscsupplyplatformserviceabilitydeleteAPIResponse struct {
	model.CommonResponse
	AlibabasscsupplyplatformserviceabilitydeleteAPIResponseModel
}

// AlibabasscsupplyplatformserviceabilitydeleteAPIResponseModel is 删除服务能力 成功返回结果
type AlibabasscsupplyplatformserviceabilitydeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceability_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabasscsupplyplatformserviceabilitydeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
