package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReportDiagnoseTempmessageReceive 导医通报告解读临时消息接收
// alibaba.alihealth.examination.report.diagnose.tempmessage.receive
//
// 导医通报告解读临时消息接收
func AlibabaAlihealthExaminationReportDiagnoseTempmessageReceive(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIRequest, resp *examination.AlibabaAlihealthExaminationReportDiagnoseTempmessageReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
