package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdReportQueryKeywordEffect 关键词报告
// alibaba.scbp.ad.report.query.keyword.effect
//
// 关键词报告
func AlibabaScbpAdReportQueryKeywordEffect(clt *core.SDKClient, req *scbp.AlibabaScbpAdReportQueryKeywordEffectAPIRequest, resp *scbp.AlibabaScbpAdReportQueryKeywordEffectAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
