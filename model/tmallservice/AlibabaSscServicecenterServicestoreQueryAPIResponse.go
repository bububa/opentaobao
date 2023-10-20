package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscservicecenterservicestorequeryAPIResponse 根据天猫id查询门店信息 API返回值
// alibaba.ssc.servicecenter.servicestore.query
//
// 根据天猫id查询门店信息
type AlibabasscservicecenterservicestorequeryAPIResponse struct {
	model.CommonResponse
	AlibabasscservicecenterservicestorequeryAPIResponseModel
}

// AlibabasscservicecenterservicestorequeryAPIResponseModel is 根据天猫id查询门店信息 成功返回结果
type AlibabasscservicecenterservicestorequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_servicecenter_servicestore_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabasscservicecenterservicestorequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
