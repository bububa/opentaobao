package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugtracetoplsydservicegetenddateAPIResponse 获取企业服务截止时间 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.service.getenddate
//
// 获取企业服务截止时间
type AlibabaalihealthdrugtracetoplsydservicegetenddateAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugtracetoplsydservicegetenddateAPIResponseModel
}

// AlibabaalihealthdrugtracetoplsydservicegetenddateAPIResponseModel is 获取企业服务截止时间 成功返回结果
type AlibabaalihealthdrugtracetoplsydservicegetenddateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_service_getenddate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回
	Result *AlibabaalihealthdrugtracetoplsydservicegetenddateResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
