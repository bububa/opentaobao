package kclub

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKclubKcQaSearchAPIRequest 知识云-知识检索 API请求
// alibaba.kclub.kc.qa.search
//
// 知识云-知识搜索服务
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaKclubKcQaSearchAPIRequest) Reset() {
	r._query = nil
	r._auth = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaKclubKcQaSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.kclub.kc.qa.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaKclubKcQaSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaKclubKcQaSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询参数
func (r *AlibabaKclubKcQaSearchAPIRequest) SetQuery(_query *KcSearchQuestionQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaKclubKcQaSearchAPIRequest) GetQuery() *KcSearchQuestionQuery {
	return r._query
}

// SetAuth is Auth Setter
// 鉴权
func (r *AlibabaKclubKcQaSearchAPIRequest) SetAuth(_auth *TenancyAuth) error {
	r._auth = _auth
	r.Set("auth", _auth)
	return nil
}

// GetAuth Auth Getter
func (r AlibabaKclubKcQaSearchAPIRequest) GetAuth() *TenancyAuth {
	return r._auth
}

var poolAlibabaKclubKcQaSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaKclubKcQaSearchRequest()
	},
}

// GetAlibabaKclubKcQaSearchRequest 从 sync.Pool 获取 AlibabaKclubKcQaSearchAPIRequest
func GetAlibabaKclubKcQaSearchAPIRequest() *AlibabaKclubKcQaSearchAPIRequest {
	return poolAlibabaKclubKcQaSearchAPIRequest.Get().(*AlibabaKclubKcQaSearchAPIRequest)
}

// ReleaseAlibabaKclubKcQaSearchAPIRequest 将 AlibabaKclubKcQaSearchAPIRequest 放入 sync.Pool
func ReleaseAlibabaKclubKcQaSearchAPIRequest(v *AlibabaKclubKcQaSearchAPIRequest) {
	v.Reset()
	poolAlibabaKclubKcQaSearchAPIRequest.Put(v)
}
