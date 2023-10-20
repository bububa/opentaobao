package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIResponse 工人请假 API返回值
// alibaba.ssc.supplyplatform.serviceworker.wokrerleave
//
// 工人请假
type AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIResponseModel
}

// AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIResponseModel is 工人请假 成功返回结果
type AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceworker_wokrerleave_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServiceworkerWokrerleaveResult `json:"result,omitempty" xml:"result,omitempty"`
}
