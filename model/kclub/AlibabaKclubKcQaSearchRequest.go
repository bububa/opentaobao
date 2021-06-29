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
    _query   *KcSearchQuestionQuery
    // 鉴权
    _auth   *TenancyAuth
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
func (r *AlibabaKclubKcQaSearchRequest) SetQuery(_query *KcSearchQuestionQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaKclubKcQaSearchRequest) GetQuery() *KcSearchQuestionQuery {
    return r._query
}
// Auth Setter
// 鉴权
func (r *AlibabaKclubKcQaSearchRequest) SetAuth(_auth *TenancyAuth) error {
    r._auth = _auth
    r.Set("auth", _auth)
    return nil
}

// Auth Getter
func (r AlibabaKclubKcQaSearchRequest) GetAuth() *TenancyAuth {
    return r._auth
}
