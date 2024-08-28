package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdReportGetTargetReport 定向报告
// alibaba.scbp.ad.report.get.target.report
//
// 定向报告
func AlibabaScbpAdReportGetTargetReport(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdReportGetTargetReportAPIRequest, resp *scbp.AlibabaScbpAdReportGetTargetReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
