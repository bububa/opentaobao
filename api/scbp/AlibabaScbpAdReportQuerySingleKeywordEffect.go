package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

/* AlibabaScbpAdReportQuerySingleKeywordEffect
单个关键词报告
alibaba.scbp.ad.report.query.single.keyword.effect

单个关键词报告 */
func AlibabaScbpAdReportQuerySingleKeywordEffect(clt *core.SDKClient, req *scbp.AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest, session string) (*scbp.AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse, error) {
	var resp scbp.AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
