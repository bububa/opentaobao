package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个关键词效果报表 API请求
alibaba.scbp.effect.keyword.single.get

单个关键词效果报表
*/
type AlibabaScbpEffectKeywordSingleGetAPIRequest struct {
    model.Params
    // IKeywordQuery
    _p4pKeywordReportQuery   *IKeywordQuery
}

// 初始化AlibabaScbpEffectKeywordSingleGetAPIRequest对象
func NewAlibabaScbpEffectKeywordSingleGetRequest() *AlibabaScbpEffectKeywordSingleGetAPIRequest{
    return &AlibabaScbpEffectKeywordSingleGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectKeywordSingleGetAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.keyword.single.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectKeywordSingleGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// P4pKeywordReportQuery Setter
// IKeywordQuery
func (r *AlibabaScbpEffectKeywordSingleGetAPIRequest) SetP4pKeywordReportQuery(_p4pKeywordReportQuery *IKeywordQuery) error {
    r._p4pKeywordReportQuery = _p4pKeywordReportQuery
    r.Set("p4p_keyword_report_query", _p4pKeywordReportQuery)
    return nil
}

// P4pKeywordReportQuery Getter
func (r AlibabaScbpEffectKeywordSingleGetAPIRequest) GetP4pKeywordReportQuery() *IKeywordQuery {
    return r._p4pKeywordReportQuery
}
