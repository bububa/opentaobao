package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordrankpricebatchgetAPIRequest 外贸直通车关键词前五名批量排价 API请求
// alibaba.scbp.ad.keyword.rank.price.batchget
//
// 外贸直通车关键词前五名批量排价
type AlibabascbpadkeywordrankpricebatchgetAPIRequest struct {
	model.Params
	// keyword_request
	_keywordRequest *TopKeywordListDto
	// 上下文
	_context *ContextDto
}

// NewAlibabascbpadkeywordrankpricebatchgetRequest 初始化AlibabascbpadkeywordrankpricebatchgetAPIRequest对象
func NewAlibabascbpadkeywordrankpricebatchgetRequest() *AlibabascbpadkeywordrankpricebatchgetAPIRequest {
	return &AlibabascbpadkeywordrankpricebatchgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordrankpricebatchgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.rank.price.batchget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordrankpricebatchgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordrankpricebatchgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeywordRequest is KeywordRequest Setter
// keyword_request
func (r *AlibabascbpadkeywordrankpricebatchgetAPIRequest) SetKeywordRequest(_keywordRequest *TopKeywordListDto) error {
	r._keywordRequest = _keywordRequest
	r.Set("keyword_request", _keywordRequest)
	return nil
}

// GetKeywordRequest KeywordRequest Getter
func (r AlibabascbpadkeywordrankpricebatchgetAPIRequest) GetKeywordRequest() *TopKeywordListDto {
	return r._keywordRequest
}

// SetContext is Context Setter
// 上下文
func (r *AlibabascbpadkeywordrankpricebatchgetAPIRequest) SetContext(_context *ContextDto) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabascbpadkeywordrankpricebatchgetAPIRequest) GetContext() *ContextDto {
	return r._context
}
