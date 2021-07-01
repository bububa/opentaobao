package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdKeywordGetAPIRequest
外贸直通车查询关键词 API请求
alibaba.scbp.ad.keyword.get

外贸直通车查询关键词 */
type AlibabaScbpAdKeywordGetAPIRequest struct {
	model.Params
	// KeywordQuery
	_queryDto *KeywordQuery
}

// NewAlibabaScbpAdKeywordGetRequest 初始化AlibabaScbpAdKeywordGetAPIRequest对象
func NewAlibabaScbpAdKeywordGetRequest() *AlibabaScbpAdKeywordGetAPIRequest {
	return &AlibabaScbpAdKeywordGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is QueryDto Setter
// KeywordQuery
func (r *AlibabaScbpAdKeywordGetAPIRequest) SetQueryDto(_queryDto *KeywordQuery) error {
	r._queryDto = _queryDto
	r.Set("query_dto", _queryDto)
	return nil
}

// Get QueryDto Getter
func (r AlibabaScbpAdKeywordGetAPIRequest) GetQueryDto() *KeywordQuery {
	return r._queryDto
}
