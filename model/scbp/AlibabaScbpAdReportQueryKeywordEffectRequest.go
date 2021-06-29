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
    _keywordReportOperation   *KeywordReportOperationDTO
    // 用户信息
    _topContext   *TopContextDTO
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
func (r *AlibabaScbpAdReportQueryKeywordEffectRequest) SetKeywordReportOperation(_keywordReportOperation *KeywordReportOperationDTO) error {
    r._keywordReportOperation = _keywordReportOperation
    r.Set("keyword_report_operation", _keywordReportOperation)
    return nil
}

// KeywordReportOperation Getter
func (r AlibabaScbpAdReportQueryKeywordEffectRequest) GetKeywordReportOperation() *KeywordReportOperationDTO {
    return r._keywordReportOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportQueryKeywordEffectRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdReportQueryKeywordEffectRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
