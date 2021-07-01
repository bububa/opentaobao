package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest
导医通报告解读临时消息接收 API请求
alibaba.alihealth.examination.report.diagnose.tempmessage.receive

导医通报告解读临时消息接收 */
type AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest struct {
	model.Params
	// 入参对象
	_reportDiagnoseImMessageRequest *ReportDiagnoseImMessageRequest
}

// NewAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest 初始化AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest() *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest {
	return &AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.tempmessage.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ReportDiagnoseImMessageRequest Setter
// 入参对象
func (r *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest) SetReportDiagnoseImMessageRequest(_reportDiagnoseImMessageRequest *ReportDiagnoseImMessageRequest) error {
	r._reportDiagnoseImMessageRequest = _reportDiagnoseImMessageRequest
	r.Set("report_diagnose_im_message_request", _reportDiagnoseImMessageRequest)
	return nil
}

// Get ReportDiagnoseImMessageRequest Getter
func (r AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest) GetReportDiagnoseImMessageRequest() *ReportDiagnoseImMessageRequest {
	return r._reportDiagnoseImMessageRequest
}
