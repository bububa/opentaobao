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
    query   *KcSearchQuestionQuery
    // 鉴权
    auth   *TenancyAuth
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
func (r *AlibabaKclubKcQaSearchPageRequest) SetQuery(query *KcSearchQuestionQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaKclubKcQaSearchPageRequest) GetQuery() *KcSearchQuestionQuery {
    return r.query
}
// Auth Setter
// 鉴权
func (r *AlibabaKclubKcQaSearchPageRequest) SetAuth(auth *TenancyAuth) error {
    r.auth = auth
    r.Set("auth", auth)
    return nil
}

// Auth Getter
func (r AlibabaKclubKcQaSearchPageRequest) GetAuth() *TenancyAuth {
    return r.auth
}
