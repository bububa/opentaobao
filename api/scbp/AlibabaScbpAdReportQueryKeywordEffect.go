package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdReportQueryKeywordEffect 关键词报告
// alibaba.scbp.ad.report.query.keyword.effect
//
// 关键词报告
func AlibabaScbpAdReportQueryKeywordEffect(clt *core.SDKClient, req *scbp.AlibabaScbpAdReportQueryKeywordEffectAPIRequest, session string) (*scbp.AlibabaScbpAdReportQueryKeywordEffectAPIResponse, error) {
	var resp scbp.AlibabaScbpAdReportQueryKeywordEffectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
