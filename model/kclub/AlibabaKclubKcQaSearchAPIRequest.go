package kclub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaKclubKcQaSearchAPIRequest
知识云-知识检索 API请求
alibaba.kclub.kc.qa.search

知识云-知识搜索服务 */
type AlibabaKclubKcQaSearchAPIRequest struct {
	model.Params
	// 查询参数
	_query *KcSearchQuestionQuery
	// 鉴权
	_auth *TenancyAuth
}

// NewAlibabaKclubKcQaSearchRequest 初始化AlibabaKclubKcQaSearchAPIRequest对象
func NewAlibabaKclubKcQaSearchRequest() *AlibabaKclubKcQaSearchAPIRequest {
	return &AlibabaKclubKcQaSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaKclubKcQaSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.kclub.kc.qa.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaKclubKcQaSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Query Setter
// 查询参数
func (r *AlibabaKclubKcQaSearchAPIRequest) SetQuery(_query *KcSearchQuestionQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r AlibabaKclubKcQaSearchAPIRequest) GetQuery() *KcSearchQuestionQuery {
	return r._query
}

// Set is Auth Setter
// 鉴权
func (r *AlibabaKclubKcQaSearchAPIRequest) SetAuth(_auth *TenancyAuth) error {
	r._auth = _auth
	r.Set("auth", _auth)
	return nil
}

// Get Auth Getter
func (r AlibabaKclubKcQaSearchAPIRequest) GetAuth() *TenancyAuth {
	return r._auth
}
