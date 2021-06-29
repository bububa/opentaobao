package kclub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-知识检索 API请求
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

// 初始化AlibabaKclubKcQaSearchRequest对象
func NewAlibabaKclubKcQaSearchRequest() *AlibabaKclubKcQaSearchRequest{
    return &AlibabaKclubKcQaSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaKclubKcQaSearchRequest) GetApiMethodName() string {
    return "alibaba.kclub.kc.qa.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaKclubKcQaSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 查询参数
func (r *AlibabaKclubKcQaSearchRequest) SetQuery(query *KcSearchQuestionQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaKclubKcQaSearchRequest) GetQuery() *KcSearchQuestionQuery {
    return r.query
}
// Auth Setter
// 鉴权
func (r *AlibabaKclubKcQaSearchRequest) SetAuth(auth *TenancyAuth) error {
    r.auth = auth
    r.Set("auth", auth)
    return nil
}

// Auth Getter
func (r AlibabaKclubKcQaSearchRequest) GetAuth() *TenancyAuth {
    return r.auth
}
