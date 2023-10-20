package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpreckeywordsysgetAPIRequest 系统推荐 API请求
// alibaba.scbp.reckeyword.sys.get
//
// 查询系统推荐词
type AlibabascbpreckeywordsysgetAPIRequest struct {
	model.Params
	// RecKeywordQuery
	_queryDto *RecKeywordQuery
}

// NewAlibabascbpreckeywordsysgetRequest 初始化AlibabascbpreckeywordsysgetAPIRequest对象
func NewAlibabascbpreckeywordsysgetRequest() *AlibabascbpreckeywordsysgetAPIRequest {
	return &AlibabascbpreckeywordsysgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpreckeywordsysgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.reckeyword.sys.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpreckeywordsysgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpreckeywordsysgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryDto is QueryDto Setter
// RecKeywordQuery
func (r *AlibabascbpreckeywordsysgetAPIRequest) SetQueryDto(_queryDto *RecKeywordQuery) error {
	r._queryDto = _queryDto
	r.Set("query_dto", _queryDto)
	return nil
}

// GetQueryDto QueryDto Getter
func (r AlibabascbpreckeywordsysgetAPIRequest) GetQueryDto() *RecKeywordQuery {
	return r._queryDto
}
