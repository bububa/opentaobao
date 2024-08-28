package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdReportGetAccountReport 账户报告
// alibaba.scbp.ad.report.get.account.report
//
// 账户报告
func AlibabaScbpAdReportGetAccountReport(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdReportGetAccountReportAPIRequest, resp *scbp.AlibabaScbpAdReportGetAccountReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
