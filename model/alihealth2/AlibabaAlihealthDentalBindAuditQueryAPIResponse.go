package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdentalbindauditqueryAPIResponse ISV查询绑定审核状态 API返回值
// alibaba.alihealth.dental.bind.audit.query
//
// ISV查询绑定审核状态
type AlibabaalihealthdentalbindauditqueryAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdentalbindauditqueryAPIResponseModel
}

// AlibabaalihealthdentalbindauditqueryAPIResponseModel is ISV查询绑定审核状态 成功返回结果
type AlibabaalihealthdentalbindauditqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_bind_audit_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaalihealthdentalbindauditqueryMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
