package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordDeleteAPIRequest 外贸直通车删除关键词 API请求
// alibaba.scbp.ad.keyword.delete
//
// 外贸直通车删除关键词
type AlibabaScbpAdKeywordDeleteAPIRequest struct {
	model.Params
	// 要删除的关键词
	_adKeyword string
}

// NewAlibabaScbpAdKeywordDeleteRequest 初始化AlibabaScbpAdKeywordDeleteAPIRequest对象
func NewAlibabaScbpAdKeywordDeleteRequest() *AlibabaScbpAdKeywordDeleteAPIRequest {
	return &AlibabaScbpAdKeywordDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdKeywordDeleteAPIRequest) Reset() {
	r._adKeyword = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdKeywordDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdKeyword is AdKeyword Setter
// 要删除的关键词
func (r *AlibabaScbpAdKeywordDeleteAPIRequest) SetAdKeyword(_adKeyword string) error {
	r._adKeyword = _adKeyword
	r.Set("ad_keyword", _adKeyword)
	return nil
}

// GetAdKeyword AdKeyword Getter
func (r AlibabaScbpAdKeywordDeleteAPIRequest) GetAdKeyword() string {
	return r._adKeyword
}

var poolAlibabaScbpAdKeywordDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdKeywordDeleteRequest()
	},
}

// GetAlibabaScbpAdKeywordDeleteRequest 从 sync.Pool 获取 AlibabaScbpAdKeywordDeleteAPIRequest
func GetAlibabaScbpAdKeywordDeleteAPIRequest() *AlibabaScbpAdKeywordDeleteAPIRequest {
	return poolAlibabaScbpAdKeywordDeleteAPIRequest.Get().(*AlibabaScbpAdKeywordDeleteAPIRequest)
}

// ReleaseAlibabaScbpAdKeywordDeleteAPIRequest 将 AlibabaScbpAdKeywordDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdKeywordDeleteAPIRequest(v *AlibabaScbpAdKeywordDeleteAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdKeywordDeleteAPIRequest.Put(v)
}
