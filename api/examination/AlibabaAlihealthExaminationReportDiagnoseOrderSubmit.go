package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderSubmit 体检报告人工解读订单
// alibaba.alihealth.examination.report.diagnose.order.submit
//
// 体检报告人工解读订单信息推送给ISV，进行人工解读
func AlibabaAlihealthExaminationReportDiagnoseOrderSubmit(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIRequest, resp *examination.AlibabaAlihealthExaminationReportDiagnoseOrderSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
