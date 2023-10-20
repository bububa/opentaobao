package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdReportGetProductReport 产品报告
// alibaba.scbp.ad.report.get.product.report
//
// 产品报告
func AlibabaScbpAdReportGetProductReport(clt *core.SDKClient, req *scbp.AlibabaScbpAdReportGetProductReportAPIRequest, resp *scbp.AlibabaScbpAdReportGetProductReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
