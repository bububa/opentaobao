package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest
报告解读订单状态更新 API请求
alibaba.alihealth.examination.report.diagnose.order.status

报告解读订单状态更新 */
type AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest struct {
	model.Params
	// 参数对象
	_reportOrderStatusRequest *ReportOrderStatusRequest
}

// New
