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
type AlibabaScbpEffectKeywordSingleGetRequest struct {
    model.Params
    // IKeywordQuery
    p4pKeywordReportQuery   *IKeywordQuery
}

// 初始化AlibabaScbpEffectKeywordSingleGetRequest对象
func NewAlibabaScbpEffectKeywordSingleGetRequest() *AlibabaScbpEffectKeywordSingleGetRequest{
    return &AlibabaScbpEffectKeywordSingleGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectKeywordSingleGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.keyword.single.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectKeywordSingleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// P4pKeywordReportQuery Setter
// IKeywordQuery
func (r *AlibabaScbpEffectKeywordSingleGetRequest) SetP4pKeywordReportQuery(p4pKeywordReportQuery *IKeywordQuery) error {
    r.p4pKeywordReportQuery = p4pKeywordReportQuery
    r.Set("p4p_keyword_report_query", p4pKeywordReportQuery)
    return nil
}

// P4pKeywordReportQuery Getter
func (r AlibabaScbpEffectKeywordSingleGetRequest) GetP4pKeywordReportQuery() *IKeywordQuery {
    return r.p4pKeywordReportQuery
}
