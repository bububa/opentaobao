package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReserveReport 体检机构对接_体检报告查询
// alibaba.alihealth.examination.reserve.report
//
// 体检机构对接_体检报告获取
func AlibabaAlihealthExaminationReserveReport(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReserveReportAPIRequest, resp *examination.AlibabaAlihealthExaminationReserveReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
