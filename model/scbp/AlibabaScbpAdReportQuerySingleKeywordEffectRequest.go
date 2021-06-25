package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个关键词报告 APIRequest
alibaba.scbp.ad.report.query.single.keyword.effect

单个关键词报告
*/
type AlibabaScbpAdReportQuerySingleKeywordEffectRequest struct {
    model.Params

    // 返回详情
    keywordReportOperation   *KeywordReportOperationDto 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdReportQuerySingleKeywordEffectRequest() *AlibabaScbpAdReportQuerySingleKeywordEffectRequest{
    return &AlibabaScbpAdReportQuerySingleKeywordEffectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdReportQuerySingleKeywordEffectRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.query.single.keyword.effect"
}

func (r AlibabaScbpAdReportQuerySingleKeywordEffectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdReportQuerySingleKeywordEffectRequest) SetKeywordReportOperation(keywordReportOperation *KeywordReportOperationDto) error {
    r.keywordReportOperation = keywordReportOperation
    r.Set("keyword_report_operation", keywordReportOperation)
    return nil
}

func (r AlibabaScbpAdReportQuerySingleKeywordEffectRequest) GetKeywordReportOperation() *KeywordReportOperationDto {
    return r.keywordReportOperation
}

func (r *AlibabaScbpAdReportQuerySingleKeywordEffectRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdReportQuerySingleKeywordEffectRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

