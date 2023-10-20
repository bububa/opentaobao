package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse 异步开方处方查询 API返回值
// alibaba.alihealth.asyncprescribe.prescription.search
//
// 异步开方处方查询
type AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponseModel
}

// AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponseModel is 异步开方处方查询 成功返回结果
type AlibabaAlihealthAsyncprescribePrescriptionSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_asyncprescribe_prescription_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
