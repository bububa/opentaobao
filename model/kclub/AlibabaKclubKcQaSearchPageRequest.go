package kclub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-知识检索(分页) API请求
alibaba.kclub.kc.qa.search.page

知识云-知识搜索服务
*/
type AlibabaKclubKcQaSearchPageRequest struct {
    model.Params
    // 查询参数
    _query   *KcSearchQuestionQuery
    // 鉴权
    _auth   *TenancyAuth
}

// 初始化AlibabaKclubKcQaSearchPageRequest对象
func NewAlibabaKclubKcQaSearchPageRequest() *AlibabaKclubKcQaSearchPageRequest{
    return &AlibabaKclubKcQaSearchPageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaKclubKcQaSearchPageRequest) GetApiMethodName() string {
    return "alibaba.kclub.kc.qa.search.page"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaKclubKcQaSearchPageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 查询参数
func (r *AlibabaKclubKcQaSearchPageRequest) SetQuery(_query *KcSearchQuestionQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaKclubKcQaSearchPageRequest) GetQuery() *KcSearchQuestionQuery {
    return r._query
}
// Auth Setter
// 鉴权
func (r *AlibabaKclubKcQaSearchPageRequest) SetAuth(_auth *TenancyAuth) error {
    r._auth = _auth
    r.Set("auth", _auth)
    return nil
}

// Auth Getter
func (r AlibabaKclubKcQaSearchPageRequest) GetAuth() *TenancyAuth {
    return r._auth
}
