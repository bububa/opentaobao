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
type AlibabaScbpAdReportQueryKeywordEffectAPIRequest struct {
    model.Params
    // 请求参数
    _keywordReportOperation   *KeywordReportOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
}

// 初始化AlibabaScbpAdReportQueryKeywordEffectAPIRequest对象
func NewAlibabaScbpAdReportQueryKeywordEffectRequest() *AlibabaScbpAdReportQueryKeywordEffectAPIRequest{
    return &AlibabaScbpAdReportQueryKeywordEffectAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.report.query.keyword.effect"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeywordReportOperation Setter
// 请求参数
func (r *AlibabaScbpAdReportQueryKeywordEffectAPIRequest) SetKeywordReportOperation(_keywordReportOperation *KeywordReportOperationDTO) error {
    r._keywordReportOperation = _keywordReportOperation
    r.Set("keyword_report_operation", _keywordReportOperation)
    return nil
}

// KeywordReportOperation Getter
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetKeywordReportOperation() *KeywordReportOperationDTO {
    return r._keywordReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportQueryKeywordEffectAPIRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportQueryKeywordEffectAPIRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
