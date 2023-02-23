package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest 外贸直通车关键词前五名批量排价 API请求
// alibaba.scbp.ad.keyword.rank.price.batchget
//
// 外贸直通车关键词前五名批量排价
type AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest struct {
	model.Params
	// keyword_request
	_keywordRequest *TopKeywordListDto
	// 上下文
	_context *ContextDto
}

// NewAlibabaScbpAdKeywordRankPriceBatchgetRequest 初始化AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest对象
func NewAlibabaScbpAdKeywordRankPriceBatchgetRequest() *AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest {
	return &AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.rank.price.batchget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeywordRequest is KeywordRequest Setter
// keyword_request
func (r *AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest) SetKeywordRequest(_keywordRequest *TopKeywordListDto) error {
	r._keywordRequest = _keywordRequest
	r.Set("keyword_request", _keywordRequest)
	return nil
}

// GetKeywordRequest KeywordRequest Getter
func (r AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest) GetKeywordRequest() *TopKeywordListDto {
	return r._keywordRequest
}

// SetContext is Context Setter
// 上下文
func (r *AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest) SetContext(_context *ContextDto) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AlibabaScbpAdKeywordRankPriceBatchgetAPIRequest) GetContext() *ContextDto {
	return r._context
}
