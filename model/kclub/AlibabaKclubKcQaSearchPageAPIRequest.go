package kclub

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKclubKcQaSearchPageAPIRequest 知识云-知识检索(分页) API请求
// alibaba.kclub.kc.qa.search.page
//
// 知识云-知识搜索服务
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaKclubKcQaSearchPageAPIRequest) Reset() {
	r._query = nil
	r._auth = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaKclubKcQaSearchPageAPIRequest) GetApiMethodName() string {
	return "alibaba.kclub.kc.qa.search.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaKclubKcQaSearchPageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaKclubKcQaSearchPageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询参数
func (r *AlibabaKclubKcQaSearchPageAPIRequest) SetQuery(_query *KcSearchQuestionQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaKclubKcQaSearchPageAPIRequest) GetQuery() *KcSearchQuestionQuery {
	return r._query
}

// SetAuth is Auth Setter
// 鉴权
func (r *AlibabaKclubKcQaSearchPageAPIRequest) SetAuth(_auth *TenancyAuth) error {
	r._auth = _auth
	r.Set("auth", _auth)
	return nil
}

// GetAuth Auth Getter
func (r AlibabaKclubKcQaSearchPageAPIRequest) GetAuth() *TenancyAuth {
	return r._auth
}

var poolAlibabaKclubKcQaSearchPageAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaKclubKcQaSearchPageRequest()
	},
}

// GetAlibabaKclubKcQaSearchPageRequest 从 sync.Pool 获取 AlibabaKclubKcQaSearchPageAPIRequest
func GetAlibabaKclubKcQaSearchPageAPIRequest() *AlibabaKclubKcQaSearchPageAPIRequest {
	return poolAlibabaKclubKcQaSearchPageAPIRequest.Get().(*AlibabaKclubKcQaSearchPageAPIRequest)
}

// ReleaseAlibabaKclubKcQaSearchPageAPIRequest 将 AlibabaKclubKcQaSearchPageAPIRequest 放入 sync.Pool
func ReleaseAlibabaKclubKcQaSearchPageAPIRequest(v *AlibabaKclubKcQaSearchPageAPIRequest) {
	v.Reset()
	poolAlibabaKclubKcQaSearchPageAPIRequest.Put(v)
}
