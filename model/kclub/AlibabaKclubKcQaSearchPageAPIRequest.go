package kclub

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaKclubKcQaSearchPageAPIRequest
知识云-知识检索(分页) API请求
alibaba.kclub.kc.qa.search.page

知识云-知识搜索服务 */
type AlibabaKclubKcQaSearchPageAPIRequest struct {
	model.Params
	// 查询参数
	_query *KcSearchQuestionQuery
	// 鉴权
	_auth *TenancyAuth
}

// NewAlibabaKclubKcQaSearchPageRequest 初始化AlibabaKclubKcQaSearchPageAPIRequest对象
func NewAlibabaKclubKcQaSearchPageRequest() *AlibabaKclubKcQaSearchPageAPIRequest {
	return &AlibabaKclubKcQaSearchPageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaKclubKcQaSearchPageAPIRequest) GetApiMethodName() string {
	return "alibaba.kclub.kc.qa.search.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaKclubKcQaSearchPageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Query Setter
// 查询参数
func (r *AlibabaKclubKcQaSearchPageAPIRequest) SetQuery(_query *KcSearchQuestionQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r AlibabaKclubKcQaSearchPageAPIRequest) GetQuery() *KcSearchQuestionQuery {
	return r._query
}

// Set is Auth Setter
// 鉴权
func (r *AlibabaKclubKcQaSearchPageAPIRequest) SetAuth(_auth *TenancyAuth) error {
	r._auth = _auth
	r.Set("auth", _auth)
	return nil
}

// Get Auth Getter
func (r AlibabaKclubKcQaSearchPageAPIRequest) GetAuth() *TenancyAuth {
	return r._auth
}
