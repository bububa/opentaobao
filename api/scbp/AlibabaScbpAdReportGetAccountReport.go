package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdReportGetAccountReport 账户报告
// alibaba.scbp.ad.report.get.account.report
//
// 账户报告
func AlibabaScbpAdReportGetAccountReport(clt *core.SDKClient, req *scbp.AlibabaScbpAdReportGetAccountReportAPIRequest, resp *scbp.AlibabaScbpAdReportGetAccountReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
