package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse 体检报告人工解读订单 API返回值
// alibaba.alihealth.examination.report.diagnose.order.submit
//
// 体检报告人工解读订单信息推送给ISV，进行人工解读
type AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponseModel is 体检报告人工解读订单 成功返回结果
type AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_order_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求状态
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 三方订单信息
	OrderInfo *OrderInfo `json:"order_info,omitempty" xml:"order_info,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResponseCode = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.OrderInfo = nil
}

var poolAlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse
func GetAlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse() *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse {
	return poolAlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse.Get().(*AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse 将 AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse(v *AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse.Put(v)
}
