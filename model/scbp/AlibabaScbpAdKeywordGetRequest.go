package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车查询关键词 APIRequest
alibaba.scbp.ad.keyword.get

外贸直通车查询关键词
*/
type AlibabaScbpAdKeywordGetRequest struct {
    model.Params

    // KeywordQuery
    queryDto   *KeywordQuery 

}

func NewAlibabaScbpAdKeywordGetRequest() *AlibabaScbpAdKeywordGetRequest{
    return &AlibabaScbpAdKeywordGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdKeywordGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.keyword.get"
}

func (r AlibabaScbpAdKeywordGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdKeywordGetRequest) SetQueryDto(queryDto *KeywordQuery) error {
    r.queryDto = queryDto
    r.Set("query_dto", queryDto)
    return nil
}

func (r AlibabaScbpAdKeywordGetRequest) GetQueryDto() *KeywordQuery {
    return r.queryDto
}

