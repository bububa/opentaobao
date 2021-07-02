package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpReckeywordSearchAPIRequest 推荐词-词推词 API请求
// alibaba.scbp.reckeyword.search
//
// 推荐词-词推词
type AlibabaScbpReckeywordSearchAPIRequest struct {
	model.Params
	// RecKeywordQuery
	_queryDto *RecKeywordQuery
}

// NewAlibabaScbpReckeywordSearchRequest 初始化AlibabaScbpReckeywordSearchAPIRequest对象
func NewAlibabaScbpReckeywordSearchRequest() *AlibabaScbpReckeywordSearchAPIRequest {
	return &AlibabaScbpReckeywordSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpReckeywordSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.reckeyword.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpReckeywordSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQueryDto is QueryDto Setter
// RecKeywordQuery
func (r *AlibabaScbpReckeywordSearchAPIRequest) SetQueryDto(_queryDto *RecKeywordQuery) error {
	r._queryDto = _queryDto
	r.Set("query_dto", _queryDto)
	return nil
}

// GetQueryDto QueryDto Getter
func (r AlibabaScbpReckeywordSearchAPIRequest) GetQueryDto() *RecKeywordQuery {
	return r._queryDto
}
