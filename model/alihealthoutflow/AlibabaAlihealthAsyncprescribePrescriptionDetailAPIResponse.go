package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthasyncprescribeprescriptiondetailAPIResponse 异步开方处方详情 API返回值
// alibaba.alihealth.asyncprescribe.prescription.detail
//
// 异步开方处方查询
type AlibabaalihealthasyncprescribeprescriptiondetailAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthasyncprescribeprescriptiondetailAPIResponseModel
}

// AlibabaalihealthasyncprescribeprescriptiondetailAPIResponseModel is 异步开方处方详情 成功返回结果
type AlibabaalihealthasyncprescribeprescriptiondetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_asyncprescribe_prescription_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
