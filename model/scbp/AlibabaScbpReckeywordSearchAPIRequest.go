package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpreckeywordsearchAPIRequest 推荐词-词推词 API请求
// alibaba.scbp.reckeyword.search
//
// 推荐词-词推词
type AlibabascbpreckeywordsearchAPIRequest struct {
	model.Params
	// RecKeywordQuery
	_queryDto *RecKeywordQuery
}

// NewAlibabascbpreckeywordsearchRequest 初始化AlibabascbpreckeywordsearchAPIRequest对象
func NewAlibabascbpreckeywordsearchRequest() *AlibabascbpreckeywordsearchAPIRequest {
	return &AlibabascbpreckeywordsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpreckeywordsearchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.reckeyword.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpreckeywordsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpreckeywordsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryDto is QueryDto Setter
// RecKeywordQuery
func (r *AlibabascbpreckeywordsearchAPIRequest) SetQueryDto(_queryDto *RecKeywordQuery) error {
	r._queryDto = _queryDto
	r.Set("query_dto", _queryDto)
	return nil
}

// GetQueryDto QueryDto Getter
func (r AlibabascbpreckeywordsearchAPIRequest) GetQueryDto() *RecKeywordQuery {
	return r._queryDto
}
