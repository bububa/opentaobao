package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIRequest 导医通报告解读临时消息接收 API请求
// alibaba.alihealth.examination.report.diagnose.tempmessage.receive
//
// 导医通报告解读临时消息接收
type AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIRequest struct {
	model.Params
	// 入参对象
	_reportDiagnoseImMessageRequest *ReportDiagnoseImMessageRequest
}

// NewAlibabaalihealthexaminationreportdiagnosetempmessagereceiveRequest 初始化AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIRequest对象
func NewAlibabaalihealthexaminationreportdiagnosetempmessagereceiveRequest() *AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIRequest {
	return &AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.tempmessage.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportDiagnoseImMessageRequest is ReportDiagnoseImMessageRequest Setter
// 入参对象
func (r *AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIRequest) SetReportDiagnoseImMessageRequest(_reportDiagnoseImMessageRequest *ReportDiagnoseImMessageRequest) error {
	r._reportDiagnoseImMessageRequest = _reportDiagnoseImMessageRequest
	r.Set("report_diagnose_im_message_request", _reportDiagnoseImMessageRequest)
	return nil
}

// GetReportDiagnoseImMessageRequest ReportDiagnoseImMessageRequest Getter
func (r AlibabaalihealthexaminationreportdiagnosetempmessagereceiveAPIRequest) GetReportDiagnoseImMessageRequest() *ReportDiagnoseImMessageRequest {
	return r._reportDiagnoseImMessageRequest
}
