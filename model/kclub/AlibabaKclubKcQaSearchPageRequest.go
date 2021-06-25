package kclub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-知识检索(分页) APIRequest
alibaba.kclub.kc.qa.search.page

知识云-知识搜索服务
*/
type AlibabaKclubKcQaSearchPageRequest struct {
    model.Params

    // 查询参数
    query   *KcSearchQuestionQuery 

    // 鉴权
    auth   *TenancyAuth 

}

func NewAlibabaKclubKcQaSearchPageRequest() *AlibabaKclubKcQaSearchPageRequest{
    return &AlibabaKclubKcQaSearchPageRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaKclubKcQaSearchPageRequest) GetApiMethodName() string {
    return "alibaba.kclub.kc.qa.search.page"
}

func (r AlibabaKclubKcQaSearchPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaKclubKcQaSearchPageRequest) SetQuery(query *KcSearchQuestionQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaKclubKcQaSearchPageRequest) GetQuery() *KcSearchQuestionQuery {
    return r.query
}

func (r *AlibabaKclubKcQaSearchPageRequest) SetAuth(auth *TenancyAuth) error {
    r.auth = auth
    r.Set("auth", auth)
    return nil
}

func (r AlibabaKclubKcQaSearchPageRequest) GetAuth() *TenancyAuth {
    return r.auth
}

