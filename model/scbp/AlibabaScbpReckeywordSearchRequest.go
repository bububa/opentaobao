package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推荐词-词推词 APIRequest
alibaba.scbp.reckeyword.search

推荐词-词推词
*/
type AlibabaScbpReckeywordSearchRequest struct {
    model.Params

    // RecKeywordQuery
    queryDto   *RecKeywordQuery 

}

func NewAlibabaScbpReckeywordSearchRequest() *AlibabaScbpReckeywordSearchRequest{
    return &AlibabaScbpReckeywordSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpReckeywordSearchRequest) GetApiMethodName() string {
    return "alibaba.scbp.reckeyword.search"
}

func (r AlibabaScbpReckeywordSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpReckeywordSearchRequest) SetQueryDto(queryDto *RecKeywordQuery) error {
    r.queryDto = queryDto
    r.Set("query_dto", queryDto)
    return nil
}

func (r AlibabaScbpReckeywordSearchRequest) GetQueryDto() *RecKeywordQuery {
    return r.queryDto
}

