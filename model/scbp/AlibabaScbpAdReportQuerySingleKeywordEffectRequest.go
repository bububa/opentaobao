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
type AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest struct {
    model.Params
    // 返回详情
    _keywordReportOperation   *KeywordReportOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest对象
func NewAlibabaScbpAdReportQuerySingleKeywordEffectRequest() *AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest{
    return &AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.query.single.keyword.effect"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeywordReportOperation Setter
// 返回详情
func (r *AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) SetKeywordReportOperation(_keywordReportOperation *KeywordReportOperationDTO) error {
    r._keywordReportOperation = _keywordReportOperation
    r.Set("keyword_report_operation", _keywordReportOperation)
    return nil
}

// KeywordReportOperation Getter
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetKeywordReportOperation() *KeywordReportOperationDTO {
    return r._keywordReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
