package kclub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabakclubkcqueryknowledgeAPIRequest 知识云-通用知识查询服务 API请求
// alibaba.kclub.kc.queryknowledge
//
// 知识云-通用知识查询服务。通过租户id、类目id、知识类型、知识状态等条件查询类目。
type AlibabakclubkcqueryknowledgeAPIRequest struct {
	model.Params
	// 查询条件
	_kcQaQuery *KcQaQuery
	// 鉴权
	_auth *TenancyAuth
}

// NewAlibabakclubkcqueryknowledgeRequest 初始化AlibabakclubkcqueryknowledgeAPIRequest对象
func NewAlibabakclubkcqueryknowledgeRequest() *AlibabakclubkcqueryknowledgeAPIRequest {
	return &AlibabakclubkcqueryknowledgeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabakclubkcqueryknowledgeAPIRequest) GetApiMethodName() string {
	return "alibaba.kclub.kc.queryknowledge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabakclubkcqueryknowledgeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabakclubkcqueryknowledgeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKcQaQuery is KcQaQuery Setter
// 查询条件
func (r *AlibabakclubkcqueryknowledgeAPIRequest) SetKcQaQuery(_kcQaQuery *KcQaQuery) error {
	r._kcQaQuery = _kcQaQuery
	r.Set("kc_qa_query", _kcQaQuery)
	return nil
}

// GetKcQaQuery KcQaQuery Getter
func (r AlibabakclubkcqueryknowledgeAPIRequest) GetKcQaQuery() *KcQaQuery {
	return r._kcQaQuery
}

// SetAuth is Auth Setter
// 鉴权
func (r *AlibabakclubkcqueryknowledgeAPIRequest) SetAuth(_auth *TenancyAuth) error {
	r._auth = _auth
	r.Set("auth", _auth)
	return nil
}

// GetAuth Auth Getter
func (r AlibabakclubkcqueryknowledgeAPIRequest) GetAuth() *TenancyAuth {
	return r._auth
}
