package kclub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-查询单个知识详情 APIRequest
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

func NewAlibabaKclubKcQaGetRequest() *AlibabaKclubKcQaGetRequest{
    return &AlibabaKclubKcQaGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaKclubKcQaGetRequest) GetApiMethodName() string {
    return "alibaba.kclub.kc.qa.get"
}

func (r AlibabaKclubKcQaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaKclubKcQaGetRequest) SetQuestionId(questionId int64) error {
    r.questionId = questionId
    r.Set("question_id", questionId)
    return nil
}

func (r AlibabaKclubKcQaGetRequest) GetQuestionId() int64 {
    return r.questionId
}

func (r *AlibabaKclubKcQaGetRequest) SetFilter(filter *KcQaFilter) error {
    r.filter = filter
    r.Set("filter", filter)
    return nil
}

func (r AlibabaKclubKcQaGetRequest) GetFilter() *KcQaFilter {
    return r.filter
}

func (r *AlibabaKclubKcQaGetRequest) SetAuth(auth *TenancyAuth) error {
    r.auth = auth
    r.Set("auth", auth)
    return nil
}

func (r AlibabaKclubKcQaGetRequest) GetAuth() *TenancyAuth {
    return r.auth
}

