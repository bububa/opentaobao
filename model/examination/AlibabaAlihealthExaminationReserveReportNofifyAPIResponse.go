package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveReportNofifyAPIResponse 服务商主动通知体检报告 API返回值
// alibaba.alihealth.examination.reserve.report.nofify
//
// 服务商主动回传用户的体检报告数据
type AlibabaAlihealthExaminationReserveReportNofifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReserveReportNofifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveReportNofifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReserveReportNofifyAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReserveReportNofifyAPIResponseModel is 服务商主动通知体检报告 成功返回结果
type AlibabaAlihealthExaminationReserveReportNofifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_report_nofify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 11
	EagleEyeTraceId string `json:"eagle_eye_trace_id,omitempty" xml:"eagle_eye_trace_id,omitempty"`
	// 返回数据对象
	Data *ReserveReportResponse `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveReportNofifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EagleEyeTraceId = ""
	m.Data = nil
}

var poolAlibabaAlihealthExaminationReserveReportNofifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReserveReportNofifyAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReserveReportNofifyAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReserveReportNofifyAPIResponse
func GetAlibabaAlihealthExaminationReserveReportNofifyAPIResponse() *AlibabaAlihealthExaminationReserveReportNofifyAPIResponse {
	return poolAlibabaAlihealthExaminationReserveReportNofifyAPIResponse.Get().(*AlibabaAlihealthExaminationReserveReportNofifyAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReserveReportNofifyAPIResponse 将 AlibabaAlihealthExaminationReserveReportNofifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReserveReportNofifyAPIResponse(v *AlibabaAlihealthExaminationReserveReportNofifyAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReserveReportNofifyAPIResponse.Put(v)
}
