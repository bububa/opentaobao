package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系统推荐 APIRequest
alibaba.scbp.reckeyword.sys.get

查询系统推荐词
*/
type AlibabaScbpReckeywordSysGetRequest struct {
    model.Params

    // RecKeywordQuery
    queryDto   *RecKeywordQuery 

}

func NewAlibabaScbpReckeywordSysGetRequest() *AlibabaScbpReckeywordSysGetRequest{
    return &AlibabaScbpReckeywordSysGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpReckeywordSysGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.reckeyword.sys.get"
}

func (r AlibabaScbpReckeywordSysGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpReckeywordSysGetRequest) SetQueryDto(queryDto *RecKeywordQuery) error {
    r.queryDto = queryDto
    r.Set("query_dto", queryDto)
    return nil
}

func (r AlibabaScbpReckeywordSysGetRequest) GetQueryDto() *RecKeywordQuery {
    return r.queryDto
}

