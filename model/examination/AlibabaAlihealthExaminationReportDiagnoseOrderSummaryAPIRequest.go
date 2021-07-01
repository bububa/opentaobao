package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest
体检报告人工解读总结回传 API请求
alibaba.alihealth.examination.report.diagnose.order.summary

记录体检报告人工解读总结 */
type AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest struct {
	model.Params
	// 入参对象
	_reportOrderSummaryRequest *ReportOrderSummaryRequest
}

// New
