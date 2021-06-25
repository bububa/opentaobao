package kclub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-知识检索 APIRequest
alibaba.kclub.kc.qa.search

知识云-知识搜索服务
*/
type AlibabaKclubKcQaSearchRequest struct {
    model.Params

    // 查询参数
    query   *KcSearchQuestionQuery 

    // 鉴权
    auth   *TenancyAuth 

}

func NewAlibabaKclubKcQaSearchRequest() *AlibabaKclubKcQaSearchRequest{
    return &AlibabaKclubKcQaSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaKclubKcQaSearchRequest) GetApiMethodName() string {
    return "alibaba.kclub.kc.qa.search"
}

func (r AlibabaKclubKcQaSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaKclubKcQaSearchRequest) SetQuery(query *KcSearchQuestionQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaKclubKcQaSearchRequest) GetQuery() *KcSearchQuestionQuery {
    return r.query
}

func (r *AlibabaKclubKcQaSearchRequest) SetAuth(auth *TenancyAuth) error {
    r.auth = auth
    r.Set("auth", auth)
    return nil
}

func (r AlibabaKclubKcQaSearchRequest) GetAuth() *TenancyAuth {
    return r.auth
}

