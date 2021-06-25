package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
单个关键词报告 
alibaba.scbp.ad.report.query.single.keyword.effect

单个关键词报告
*/
func AlibabaScbpAdReportQuerySingleKeywordEffect(clt *core.SDKClient, req *scbp.AlibabaScbpAdReportQuerySingleKeywordEffectRequest, session string) (*scbp.AlibabaScbpAdReportQuerySingleKeywordEffectResponse, error) {
    var resp scbp.AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
