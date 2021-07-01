package kclub

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-通用知识查询服务 API请求
alibaba.kclub.kc.queryknowledge

知识云-通用知识查询服务。通过租户id、类目id、知识类型、知识状态等条件查询类目。
*/
type AlibabaKclubKcQueryknowledgeAPIRequest struct {
    model.Params
    // 查询条件
    _kcQaQuery   *KcQaQuery
    // 鉴权
    _auth   *TenancyAuth
}

// 初始化AlibabaKclubKcQueryknowledgeAPIRequest对象
func NewAlibabaKclubKcQueryknowledgeRequest() *AlibabaKclubKcQueryknowledgeAPIRequest{
    return &AlibabaKclubKcQueryknowledgeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaKclubKcQueryknowledgeAPIRequest) GetApiMethodName() string {
    return "alibaba.kclub.kc.queryknowledge"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaKclubKcQueryknowledgeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KcQaQuery Setter
// 查询条件
func (r *AlibabaKclubKcQueryknowledgeAPIRequest) SetKcQaQuery(_kcQaQuery *KcQaQuery) error {
    r._kcQaQuery = _kcQaQuery
    r.Set("kc_qa_query", _kcQaQuery)
    return nil
}

// KcQaQuery Getter
func (r AlibabaKclubKcQueryknowledgeAPIRequest) GetKcQaQuery() *KcQaQuery {
    return r._kcQaQuery
}
// Auth Setter
// 鉴权
func (r *AlibabaKclubKcQueryknowledgeAPIRequest) SetAuth(_auth *TenancyAuth) error {
    r._auth = _auth
    r.Set("auth", _auth)
    return nil
}

// Auth Getter
func (r AlibabaKclubKcQueryknowledgeAPIRequest) GetAuth() *TenancyAuth {
    return r._auth
}
