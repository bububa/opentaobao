package examination

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest 体检报告人工解读总结回传 API请求
// alibaba.alihealth.examination.report.diagnose.order.summary
//
// 记录体检报告人工解读总结
type AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest struct {
	model.Params
	// 入参对象
	_reportOrderSummaryRequest *ReportOrderSummaryRequest
}

// NewAlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest 初始化AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest对象
func NewAlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest {
	return &AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest) Reset() {
	r._reportOrderSummaryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.report.diagnose.order.summary"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReportOrderSummaryRequest is ReportOrderSummaryRequest Setter
// 入参对象
func (r *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest) SetReportOrderSummaryRequest(_reportOrderSummaryRequest *ReportOrderSummaryRequest) error {
	r._reportOrderSummaryRequest = _reportOrderSummaryRequest
	r.Set("report_order_summary_request", _reportOrderSummaryRequest)
	return nil
}

// GetReportOrderSummaryRequest ReportOrderSummaryRequest Getter
func (r AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest) GetReportOrderSummaryRequest() *ReportOrderSummaryRequest {
	return r._reportOrderSummaryRequest
}

var poolAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest()
	},
}

// GetAlibabaAlihealthExaminationReportDiagnoseOrderSummaryRequest 从 sync.Pool 获取 AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest
func GetAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest() *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest {
	return poolAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest.Get().(*AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest)
}

// ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest 将 AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest(v *AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest.Put(v)
}
