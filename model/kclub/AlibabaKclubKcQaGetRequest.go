package kclub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-查询单个知识详情 API请求
alibaba.kclub.kc.qa.get

知识云-查询单个知识详情。通过租户id、问题id查询问题详情
*/
type AlibabaKclubKcQaGetRequest struct {
    model.Params
    // 问题id
    _questionId   int64
    // 过滤条件
    _filter   *KcQaFilter
    // 鉴权
    _auth   *TenancyAuth
}

// 初始化AlibabaKclubKcQaGetRequest对象
func NewAlibabaKclubKcQaGetRequest() *AlibabaKclubKcQaGetRequest{
    return &AlibabaKclubKcQaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaKclubKcQaGetRequest) GetApiMethodName() string {
    return "alibaba.kclub.kc.qa.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaKclubKcQaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QuestionId Setter
// 问题id
func (r *AlibabaKclubKcQaGetRequest) SetQuestionId(_questionId int64) error {
    r._questionId = _questionId
    r.Set("question_id", _questionId)
    return nil
}

// QuestionId Getter
func (r AlibabaKclubKcQaGetRequest) GetQuestionId() int64 {
    return r._questionId
}
// Filter Setter
// 过滤条件
func (r *AlibabaKclubKcQaGetRequest) SetFilter(_filter *KcQaFilter) error {
    r._filter = _filter
    r.Set("filter", _filter)
    return nil
}

// Filter Getter
func (r AlibabaKclubKcQaGetRequest) GetFilter() *KcQaFilter {
    return r._filter
}
// Auth Setter
// 鉴权
func (r *AlibabaKclubKcQaGetRequest) SetAuth(_auth *TenancyAuth) error {
    r._auth = _auth
    r.Set("auth", _auth)
    return nil
}

// Auth Getter
func (r AlibabaKclubKcQaGetRequest) GetAuth() *TenancyAuth {
    return r._auth
}
