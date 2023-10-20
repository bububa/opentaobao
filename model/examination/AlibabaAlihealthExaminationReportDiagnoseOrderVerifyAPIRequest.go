package examination

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest 报告解读令牌校验 API请求
// alibaba.alihealth.examination.report.diagnose.order.verify
//
// 报告解读令牌校验
type AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest struct {
	model.Params
	// 请求对象
	_reportImTokenStatusRequest *ReportImTokenStatusRequest
}

// NewAlibabaAlihealthExaminationReportDiagnoseOrderVerifyRequest 初始化AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseOrderVerifyRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest {
	return &AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest) Reset() {
	r._reportImTokenStatusRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportImTokenStatusRequest is ReportImTokenStatusRequest Setter
// 请求对象
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest) SetReportImTokenStatusRequest(_reportImTokenStatusRequest *ReportImTokenStatusRequest) error {
	r._reportImTokenStatusRequest = _reportImTokenStatusRequest
	r.Set("report_im_token_status_request", _reportImTokenStatusRequest)
	return nil
}

// GetReportImTokenStatusRequest ReportImTokenStatusRequest Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest) GetReportImTokenStatusRequest() *ReportImTokenStatusRequest {
	return r._reportImTokenStatusRequest
}

var poolAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthExaminationReportDiagnoseOrderVerifyRequest()
	},
}

// GetAlibabaAlihealthExaminationReportDiagnoseOrderVerifyRequest 从 sync.Pool 获取 AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest
func GetAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest {
	return poolAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest.Get().(*AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest)
}

// ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest 将 AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest(v *AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest.Put(v)
}
