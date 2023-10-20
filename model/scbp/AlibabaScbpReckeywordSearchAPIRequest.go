package scbp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpReckeywordSearchAPIRequest) Reset() {
	r._queryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpReckeywordSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.reckeyword.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpReckeywordSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpReckeywordSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaScbpReckeywordSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpReckeywordSearchRequest()
	},
}

// GetAlibabaScbpReckeywordSearchRequest 从 sync.Pool 获取 AlibabaScbpReckeywordSearchAPIRequest
func GetAlibabaScbpReckeywordSearchAPIRequest() *AlibabaScbpReckeywordSearchAPIRequest {
	return poolAlibabaScbpReckeywordSearchAPIRequest.Get().(*AlibabaScbpReckeywordSearchAPIRequest)
}

// ReleaseAlibabaScbpReckeywordSearchAPIRequest 将 AlibabaScbpReckeywordSearchAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpReckeywordSearchAPIRequest(v *AlibabaScbpReckeywordSearchAPIRequest) {
	v.Reset()
	poolAlibabaScbpReckeywordSearchAPIRequest.Put(v)
}
