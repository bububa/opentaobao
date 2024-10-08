package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdReportQueryKeywordEffect 关键词报告
// alibaba.scbp.ad.report.query.keyword.effect
//
// 关键词报告
func AlibabaScbpAdReportQueryKeywordEffect(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdReportQueryKeywordEffectAPIRequest, resp *scbp.AlibabaScbpAdReportQueryKeywordEffectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
