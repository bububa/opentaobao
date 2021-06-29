package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词报告 API请求
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

// 初始化AlibabaScbpAdReportQueryKeywordEffectRequest对象
func NewAlibabaScbpAdReportQueryKeywordEffectRequest() *AlibabaScbpAdReportQueryKeywordEffectRequest{
    return &AlibabaScbpAdReportQueryKeywordEffectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportQueryKeywordEffectRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.query.keyword.effect"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportQueryKeywordEffectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeywordReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportQueryKeywordEffectRequest) SetKeywordReportOperation(keywordReportOperation *KeywordReportOperationDto) error {
    r.keywordReportOperation = keywordReportOperation
    r.Set("keyword_report_operation", keywordReportOperation)
    return nil
}

// KeywordReportOperation Getter
func (r AlibabaScbpAdReportQueryKeywordEffectRequest) GetKeywordReportOperation() *KeywordReportOperationDto {
    return r.keywordReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportQueryKeywordEffectRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportQueryKeywordEffectRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
