package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderVerify 报告解读令牌校验
// alibaba.alihealth.examination.report.diagnose.order.verify
//
// 报告解读令牌校验
func AlibabaAlihealthExaminationReportDiagnoseOrderVerify(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIRequest, resp *examination.AlibabaAlihealthExaminationReportDiagnoseOrderVerifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
