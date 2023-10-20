package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdReportGetLastEffectDate 获取最近报表生成时间
// alibaba.scbp.ad.report.get.last.effect.date
//
// 获取最近报表生成时间
func AlibabaScbpAdReportGetLastEffectDate(clt *core.SDKClient, req *scbp.AlibabaScbpAdReportGetLastEffectDateAPIRequest, resp *scbp.AlibabaScbpAdReportGetLastEffectDateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
