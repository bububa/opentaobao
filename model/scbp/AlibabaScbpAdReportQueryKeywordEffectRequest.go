package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词报告 APIRequest
alibaba.scbp.ad.report.query.keyword.effect

关键词报告
*/
type AlibabaScbpAdReportQueryKeywordEffectRequest struct {
    model.Params

    // 请求参数
    keywordReportOperation   *KeywordReportOperationDto 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdReportQueryKeywordEffectRequest() *AlibabaScbpAdReportQueryKeywordEffectRequest{
    return &AlibabaScbpAdReportQueryKeywordEffectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdReportQueryKeywordEffectRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.query.keyword.effect"
}

func (r AlibabaScbpAdReportQueryKeywordEffectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdReportQueryKeywordEffectRequest) SetKeywordReportOperation(keywordReportOperation *KeywordReportOperationDto) error {
    r.keywordReportOperation = keywordReportOperation
    r.Set("keyword_report_operation", keywordReportOperation)
    return nil
}

func (r AlibabaScbpAdReportQueryKeywordEffectRequest) GetKeywordReportOperation() *KeywordReportOperationDto {
    return r.keywordReportOperation
}

func (r *AlibabaScbpAdReportQueryKeywordEffectRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdReportQueryKeywordEffectRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

