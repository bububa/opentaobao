package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGet 获取报告解读url
// alibaba.alihealth.examination.report.diagnose.order.diagnoseurl.get
//
// 获取报告解读url
func AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGet(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIRequest, resp *examination.AlibabaAlihealthExaminationReportDiagnoseOrderDiagnoseurlGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
