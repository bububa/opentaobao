package kclub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-通用知识查询服务 APIRequest
alibaba.kclub.kc.queryknowledge

知识云-通用知识查询服务。通过租户id、类目id、知识类型、知识状态等条件查询类目。
*/
type AlibabaKclubKcQueryknowledgeRequest struct {
    model.Params

    // 查询条件
    kcQaQuery   *KcQaQuery 

    // 鉴权
    auth   *TenancyAuth 

}

func NewAlibabaKclubKcQueryknowledgeRequest() *AlibabaKclubKcQueryknowledgeRequest{
    return &AlibabaKclubKcQueryknowledgeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaKclubKcQueryknowledgeRequest) GetApiMethodName() string {
    return "alibaba.kclub.kc.queryknowledge"
}

func (r AlibabaKclubKcQueryknowledgeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaKclubKcQueryknowledgeRequest) SetKcQaQuery(kcQaQuery *KcQaQuery) error {
    r.kcQaQuery = kcQaQuery
    r.Set("kc_qa_query", kcQaQuery)
    return nil
}

func (r AlibabaKclubKcQueryknowledgeRequest) GetKcQaQuery() *KcQaQuery {
    return r.kcQaQuery
}

func (r *AlibabaKclubKcQueryknowledgeRequest) SetAuth(auth *TenancyAuth) error {
    r.auth = auth
    r.Set("auth", auth)
    return nil
}

func (r AlibabaKclubKcQueryknowledgeRequest) GetAuth() *TenancyAuth {
    return r.auth
}

