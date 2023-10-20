package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse 工人取消请假 API返回值
// alibaba.ssc.supplyplatform.serviceworker.cancelleave
//
// 工人取消请假
type AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponseModel
}

// AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponseModel is 工人取消请假 成功返回结果
type AlibabaSscSupplyplatformServiceworkerCancelleaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceworker_cancelleave_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServiceworkerCancelleaveResult `json:"result,omitempty" xml:"result,omitempty"`
}
