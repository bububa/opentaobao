package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordgetAPIRequest 外贸直通车查询关键词 API请求
// alibaba.scbp.ad.keyword.get
//
// 外贸直通车查询关键词
type AlibabascbpadkeywordgetAPIRequest struct {
	model.Params
	// KeywordQuery
	_queryDto *KeywordQuery
}

// NewAlibabascbpadkeywordgetRequest 初始化AlibabascbpadkeywordgetAPIRequest对象
func NewAlibabascbpadkeywordgetRequest() *AlibabascbpadkeywordgetAPIRequest {
	return &AlibabascbpadkeywordgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryDto is QueryDto Setter
// KeywordQuery
func (r *AlibabascbpadkeywordgetAPIRequest) SetQueryDto(_queryDto *KeywordQuery) error {
	r._queryDto = _queryDto
	r.Set("query_dto", _queryDto)
	return nil
}

// GetQueryDto QueryDto Getter
func (r AlibabascbpadkeywordgetAPIRequest) GetQueryDto() *KeywordQuery {
	return r._queryDto
}
