package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpKeywordMatchedProductsGetAPIRequest 查询和词匹配的推广产品 API请求
// alibaba.scbp.keyword.matched.products.get
//
// 查询和词匹配的推广产品
type AlibabaScbpKeywordMatchedProductsGetAPIRequest struct {
	model.Params
	// 已购买的关键词
	_adKeyword string
}

// NewAlibabaScbpKeywordMatchedProductsGetRequest 初始化AlibabaScbpKeywordMatchedProductsGetAPIRequest对象
func NewAlibabaScbpKeywordMatchedProductsGetRequest() *AlibabaScbpKeywordMatchedProductsGetAPIRequest {
	return &AlibabaScbpKeywordMatchedProductsGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpKeywordMatchedProductsGetAPIRequest) Reset() {
	r._adKeyword = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpKeywordMatchedProductsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.keyword.matched.products.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpKeywordMatchedProductsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpKeywordMatchedProductsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdKeyword is AdKeyword Setter
// 已购买的关键词
func (r *AlibabaScbpKeywordMatchedProductsGetAPIRequest) SetAdKeyword(_adKeyword string) error {
	r._adKeyword = _adKeyword
	r.Set("ad_keyword", _adKeyword)
	return nil
}

// GetAdKeyword AdKeyword Getter
func (r AlibabaScbpKeywordMatchedProductsGetAPIRequest) GetAdKeyword() string {
	return r._adKeyword
}

var poolAlibabaScbpKeywordMatchedProductsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpKeywordMatchedProductsGetRequest()
	},
}

// GetAlibabaScbpKeywordMatchedProductsGetRequest 从 sync.Pool 获取 AlibabaScbpKeywordMatchedProductsGetAPIRequest
func GetAlibabaScbpKeywordMatchedProductsGetAPIRequest() *AlibabaScbpKeywordMatchedProductsGetAPIRequest {
	return poolAlibabaScbpKeywordMatchedProductsGetAPIRequest.Get().(*AlibabaScbpKeywordMatchedProductsGetAPIRequest)
}

// ReleaseAlibabaScbpKeywordMatchedProductsGetAPIRequest 将 AlibabaScbpKeywordMatchedProductsGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpKeywordMatchedProductsGetAPIRequest(v *AlibabaScbpKeywordMatchedProductsGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpKeywordMatchedProductsGetAPIRequest.Put(v)
}
