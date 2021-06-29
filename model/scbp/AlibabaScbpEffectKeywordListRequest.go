package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词报表 API请求
alibaba.scbp.effect.keyword.list

关键词报表
*/
type AlibabaScbpEffectKeywordListRequest struct {
    model.Params
    // IKeywordQuery
    _p4pKeywordReportQuery   *IKeywordQuery
}

// 初始化AlibabaScbpEffectKeywordListRequest对象
func NewAlibabaScbpEffectKeywordListRequest() *AlibabaScbpEffectKeywordListRequest{
    return &AlibabaScbpEffectKeywordListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectKeywordListRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.keyword.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectKeywordListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// P4pKeywordReportQuery Setter
// IKeywordQuery
func (r *AlibabaScbpEffectKeywordListRequest) SetP4pKeywordReportQuery(_p4pKeywordReportQuery *IKeywordQuery) error {
    r._p4pKeywordReportQuery = _p4pKeywordReportQuery
    r.Set("p4p_keyword_report_query", _p4pKeywordReportQuery)
    return nil
}

// P4pKeywordReportQuery Getter
func (r AlibabaScbpEffectKeywordListRequest) GetP4pKeywordReportQuery() *IKeywordQuery {
    return r._p4pKeywordReportQuery
}
