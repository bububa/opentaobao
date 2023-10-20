package kclub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabakclubkcqasearchpageAPIRequest 知识云-知识检索(分页) API请求
// alibaba.kclub.kc.qa.search.page
//
// 知识云-知识搜索服务
type AlibabakclubkcqasearchpageAPIRequest struct {
	model.Params
	// 查询参数
	_query *KcSearchQuestionQuery
	// 鉴权
	_auth *TenancyAuth
}

// NewAlibabakclubkcqasearchpageRequest 初始化AlibabakclubkcqasearchpageAPIRequest对象
func NewAlibabakclubkcqasearchpageRequest() *AlibabakclubkcqasearchpageAPIRequest {
	return &AlibabakclubkcqasearchpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabakclubkcqasearchpageAPIRequest) GetApiMethodName() string {
	return "alibaba.kclub.kc.qa.search.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabakclubkcqasearchpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabakclubkcqasearchpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询参数
func (r *AlibabakclubkcqasearchpageAPIRequest) SetQuery(_query *KcSearchQuestionQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabakclubkcqasearchpageAPIRequest) GetQuery() *KcSearchQuestionQuery {
	return r._query
}

// SetAuth is Auth Setter
// 鉴权
func (r *AlibabakclubkcqasearchpageAPIRequest) SetAuth(_auth *TenancyAuth) error {
	r._auth = _auth
	r.Set("auth", _auth)
	return nil
}

// GetAuth Auth Getter
func (r AlibabakclubkcqasearchpageAPIRequest) GetAuth() *TenancyAuth {
	return r._auth
}
