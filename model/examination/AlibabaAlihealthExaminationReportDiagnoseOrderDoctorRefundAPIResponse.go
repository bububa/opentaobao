package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse 报告解读订单-医生退款 API返回值
// alibaba.alihealth.examination.report.diagnose.order.doctor.refund
//
// 报告解读订单医生退款
type AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponseModel is 报告解读订单-医生退款 成功返回结果
type AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_report_diagnose_order_doctor_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 接口是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Data = false
	m.IsSuccess = false
}

var poolAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse
func GetAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse() *AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse {
	return poolAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse.Get().(*AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse 将 AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse(v *AlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReportDiagnoseOrderDoctorRefundAPIResponse.Put(v)
}
