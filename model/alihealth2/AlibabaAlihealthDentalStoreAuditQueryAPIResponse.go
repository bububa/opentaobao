package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdentalstoreauditqueryAPIResponse ISV查询门店审核状态 API返回值
// alibaba.alihealth.dental.store.audit.query
//
// ISV查询门店审核状态
type AlibabaalihealthdentalstoreauditqueryAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdentalstoreauditqueryAPIResponseModel
}

// AlibabaalihealthdentalstoreauditqueryAPIResponseModel is ISV查询门店审核状态 成功返回结果
type AlibabaalihealthdentalstoreauditqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_store_audit_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaalihealthdentalstoreauditqueryMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
