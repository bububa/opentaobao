package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderStatus 报告解读订单状态更新
// alibaba.alihealth.examination.report.diagnose.order.status
//
// 报告解读订单状态更新
func AlibabaAlihealthExaminationReportDiagnoseOrderStatus(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIRequest, resp *examination.AlibabaAlihealthExaminationReportDiagnoseOrderStatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
