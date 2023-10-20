package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdReportQuerySingleKeywordEffect 单个关键词报告
// alibaba.scbp.ad.report.query.single.keyword.effect
//
// 单个关键词报告
func AlibabaScbpAdReportQuerySingleKeywordEffect(clt *core.SDKClient, req *scbp.AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest, resp *scbp.AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
