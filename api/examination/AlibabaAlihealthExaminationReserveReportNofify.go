package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReserveReportNofify 服务商主动通知体检报告
// alibaba.alihealth.examination.reserve.report.nofify
//
// 服务商主动回传用户的体检报告数据
func AlibabaAlihealthExaminationReserveReportNofify(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveReportNofifyAPIRequest, resp *examination.AlibabaAlihealthExaminationReserveReportNofifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
