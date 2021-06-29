package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个关键词报告 API请求
alibaba.scbp.ad.report.query.single.keyword.effect

单个关键词报告
*/
type AlibabaScbpAdReportQuerySingleKeywordEffectRequest struct {
    model.Params
    // 返回详情
    _keywordReportOperation   *KeywordReportOperationDto
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdReportQuerySingleKeywordEffectRequest对象
func NewAlibabaScbpAdReportQuerySingleKeywordEffectRequest() *AlibabaScbpAdReportQuerySingleKeywordEffectRequest{
    return &AlibabaScbpAdReportQuerySingleKeywordEffectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportQuerySingleKeywordEffectRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.query.single.keyword.effect"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportQuerySingleKeywordEffectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeywordReportOperation Setter
// 返回详情
func (r *AlibabaScbpAdReportQuerySingleKeywordEffectRequest) SetKeywordReportOperation(_keywordReportOperation *KeywordReportOperationDto) error {
    r._keywordReportOperation = _keywordReportOperation
    r.Set("keyword_report_operation", _keywordReportOperation)
    return nil
}

// KeywordReportOperation Getter
func (r AlibabaScbpAdReportQuerySingleKeywordEffectRequest) GetKeywordReportOperation() *KeywordReportOperationDto {
    return r._keywordReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportQuerySingleKeywordEffectRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportQuerySingleKeywordEffectRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
