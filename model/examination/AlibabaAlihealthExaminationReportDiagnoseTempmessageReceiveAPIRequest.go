package examination

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest 导医通报告解读临时消息接收 API请求
// alibaba.alihealth.examination.report.diagnose.tempmessage.receive
//
// 导医通报告解读临时消息接收
type AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest struct {
	model.Params
	// 入参对象
	_reportDiagnoseImMessageRequest *ReportDiagnoseImMessageRequest
}

// NewAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest 初始化AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest() *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest {
	return &AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest) Reset() {
	r._reportDiagnoseImMessageRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.tempmessage.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportDiagnoseImMessageRequest is ReportDiagnoseImMessageRequest Setter
// 入参对象
func (r *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest) SetReportDiagnoseImMessageRequest(_reportDiagnoseImMessageRequest *ReportDiagnoseImMessageRequest) error {
	r._reportDiagnoseImMessageRequest = _reportDiagnoseImMessageRequest
	r.Set("report_diagnose_im_message_request", _reportDiagnoseImMessageRequest)
	return nil
}

// GetReportDiagnoseImMessageRequest ReportDiagnoseImMessageRequest Getter
func (r AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest) GetReportDiagnoseImMessageRequest() *ReportDiagnoseImMessageRequest {
	return r._reportDiagnoseImMessageRequest
}

var poolAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest()
	},
}

// GetAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveRequest 从 sync.Pool 获取 AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest
func GetAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest() *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest {
	return poolAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest.Get().(*AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest)
}

// ReleaseAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest 将 AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest(v *AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest.Put(v)
}
