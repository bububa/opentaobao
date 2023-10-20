package kclub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabakclubkcqasearchAPIRequest 知识云-知识检索 API请求
// alibaba.kclub.kc.qa.search
//
// 知识云-知识搜索服务
type AlibabakclubkcqasearchAPIRequest struct {
	model.Params
	// 查询参数
	_query *KcSearchQuestionQuery
	// 鉴权
	_auth *TenancyAuth
}

// NewAlibabakclubkcqasearchRequest 初始化AlibabakclubkcqasearchAPIRequest对象
func NewAlibabakclubkcqasearchRequest() *AlibabakclubkcqasearchAPIRequest {
	return &AlibabakclubkcqasearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabakclubkcqasearchAPIRequest) GetApiMethodName() string {
	return "alibaba.kclub.kc.qa.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabakclubkcqasearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabakclubkcqasearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询参数
func (r *AlibabakclubkcqasearchAPIRequest) SetQuery(_query *KcSearchQuestionQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabakclubkcqasearchAPIRequest) GetQuery() *KcSearchQuestionQuery {
	return r._query
}

// SetAuth is Auth Setter
// 鉴权
func (r *AlibabakclubkcqasearchAPIRequest) SetAuth(_auth *TenancyAuth) error {
	r._auth = _auth
	r.Set("auth", _auth)
	return nil
}

// GetAuth Auth Getter
func (r AlibabakclubkcqasearchAPIRequest) GetAuth() *TenancyAuth {
	return r._auth
}
