package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetopyljgservicegetenddateAPIResponse 获取服务截止日期 API返回值
// alibaba.alihealth.drugtrace.top.yljg.service.getenddate
//
// 获取企业服务截止时间
type AlibabaalihealthdrugtracetopyljgservicegetenddateAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugtracetopyljgservicegetenddateAPIResponseModel
}

// AlibabaalihealthdrugtracetopyljgservicegetenddateAPIResponseModel is 获取服务截止日期 成功返回结果
type AlibabaalihealthdrugtracetopyljgservicegetenddateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_service_getenddate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaalihealthdrugtracetopyljgservicegetenddateResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
