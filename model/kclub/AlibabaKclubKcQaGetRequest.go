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
    questionId   int64
    // 过滤条件
    filter   *KcQaFilter
    // 鉴权
    auth   *TenancyAuth
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
func (r *AlibabaKclubKcQaGetRequest) SetQuestionId(questionId int64) error {
    r.questionId = questionId
    r.Set("question_id", questionId)
    return nil
}

// QuestionId Getter
func (r AlibabaKclubKcQaGetRequest) GetQuestionId() int64 {
    return r.questionId
}
// Filter Setter
// 过滤条件
func (r *AlibabaKclubKcQaGetRequest) SetFilter(filter *KcQaFilter) error {
    r.filter = filter
    r.Set("filter", filter)
    return nil
}

// Filter Getter
func (r AlibabaKclubKcQaGetRequest) GetFilter() *KcQaFilter {
    return r.filter
}
// Auth Setter
// 鉴权
func (r *AlibabaKclubKcQaGetRequest) SetAuth(auth *TenancyAuth) error {
    r.auth = auth
    r.Set("auth", auth)
    return nil
}

// Auth Getter
func (r AlibabaKclubKcQaGetRequest) GetAuth() *TenancyAuth {
    return r.auth
}
