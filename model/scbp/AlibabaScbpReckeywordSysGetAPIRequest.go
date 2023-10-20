package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpReckeywordSysGetAPIRequest 系统推荐 API请求
// alibaba.scbp.reckeyword.sys.get
//
// 查询系统推荐词
type AlibabaScbpReckeywordSysGetAPIRequest struct {
	model.Params
	// RecKeywordQuery
	_queryDto *RecKeywordQuery
}

// NewAlibabaScbpReckeywordSysGetRequest 初始化AlibabaScbpReckeywordSysGetAPIRequest对象
func NewAlibabaScbpReckeywordSysGetRequest() *AlibabaScbpReckeywordSysGetAPIRequest {
	return &AlibabaScbpReckeywordSysGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpReckeywordSysGetAPIRequest) Reset() {
	r._queryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpReckeywordSysGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.reckeyword.sys.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpReckeywordSysGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpReckeywordSysGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryDto is QueryDto Setter
// RecKeywordQuery
func (r *AlibabaScbpReckeywordSysGetAPIRequest) SetQueryDto(_queryDto *RecKeywordQuery) error {
	r._queryDto = _queryDto
	r.Set("query_dto", _queryDto)
	return nil
}

// GetQueryDto QueryDto Getter
func (r AlibabaScbpReckeywordSysGetAPIRequest) GetQueryDto() *RecKeywordQuery {
	return r._queryDto
}

var poolAlibabaScbpReckeywordSysGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpReckeywordSysGetRequest()
	},
}

// GetAlibabaScbpReckeywordSysGetRequest 从 sync.Pool 获取 AlibabaScbpReckeywordSysGetAPIRequest
func GetAlibabaScbpReckeywordSysGetAPIRequest() *AlibabaScbpReckeywordSysGetAPIRequest {
	return poolAlibabaScbpReckeywordSysGetAPIRequest.Get().(*AlibabaScbpReckeywordSysGetAPIRequest)
}

// ReleaseAlibabaScbpReckeywordSysGetAPIRequest 将 AlibabaScbpReckeywordSysGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpReckeywordSysGetAPIRequest(v *AlibabaScbpReckeywordSysGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpReckeywordSysGetAPIRequest.Put(v)
}
