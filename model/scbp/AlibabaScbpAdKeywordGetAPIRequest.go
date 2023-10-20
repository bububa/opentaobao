package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordGetAPIRequest 外贸直通车查询关键词 API请求
// alibaba.scbp.ad.keyword.get
//
// 外贸直通车查询关键词
type AlibabaScbpAdKeywordGetAPIRequest struct {
	model.Params
	// KeywordQuery
	_queryDto *KeywordQuery
}

// NewAlibabaScbpAdKeywordGetRequest 初始化AlibabaScbpAdKeywordGetAPIRequest对象
func NewAlibabaScbpAdKeywordGetRequest() *AlibabaScbpAdKeywordGetAPIRequest {
	return &AlibabaScbpAdKeywordGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdKeywordGetAPIRequest) Reset() {
	r._queryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryDto is QueryDto Setter
// KeywordQuery
func (r *AlibabaScbpAdKeywordGetAPIRequest) SetQueryDto(_queryDto *KeywordQuery) error {
	r._queryDto = _queryDto
	r.Set("query_dto", _queryDto)
	return nil
}

// GetQueryDto QueryDto Getter
func (r AlibabaScbpAdKeywordGetAPIRequest) GetQueryDto() *KeywordQuery {
	return r._queryDto
}

var poolAlibabaScbpAdKeywordGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdKeywordGetRequest()
	},
}

// GetAlibabaScbpAdKeywordGetRequest 从 sync.Pool 获取 AlibabaScbpAdKeywordGetAPIRequest
func GetAlibabaScbpAdKeywordGetAPIRequest() *AlibabaScbpAdKeywordGetAPIRequest {
	return poolAlibabaScbpAdKeywordGetAPIRequest.Get().(*AlibabaScbpAdKeywordGetAPIRequest)
}

// ReleaseAlibabaScbpAdKeywordGetAPIRequest 将 AlibabaScbpAdKeywordGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdKeywordGetAPIRequest(v *AlibabaScbpAdKeywordGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdKeywordGetAPIRequest.Put(v)
}
