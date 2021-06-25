package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词报表 APIRequest
alibaba.scbp.effect.keyword.list

关键词报表
*/
type AlibabaScbpEffectKeywordListRequest struct {
    model.Params

    // IKeywordQuery
    p4pKeywordReportQuery   *IKeywordQuery 

}

func NewAlibabaScbpEffectKeywordListRequest() *AlibabaScbpEffectKeywordListRequest{
    return &AlibabaScbpEffectKeywordListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpEffectKeywordListRequest) GetApiMethodName() string {
    return "alibaba.scbp.effect.keyword.list"
}

func (r AlibabaScbpEffectKeywordListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpEffectKeywordListRequest) SetP4pKeywordReportQuery(p4pKeywordReportQuery *IKeywordQuery) error {
    r.p4pKeywordReportQuery = p4pKeywordReportQuery
    r.Set("p4p_keyword_report_query", p4pKeywordReportQuery)
    return nil
}

func (r AlibabaScbpEffectKeywordListRequest) GetP4pKeywordReportQuery() *IKeywordQuery {
    return r.p4pKeywordReportQuery
}

