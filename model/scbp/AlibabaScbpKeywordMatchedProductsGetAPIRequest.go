package scbp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpKeywordMatchedProductsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.keyword.matched.products.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpKeywordMatchedProductsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AdKeyword Setter
// 已购买的关键词
func (r *AlibabaScbpKeywordMatchedProductsGetAPIRequest) SetAdKeyword(_adKeyword string) error {
	r._adKeyword = _adKeyword
	r.Set("ad_keyword", _adKeyword)
	return nil
}

// Get AdKeyword Getter
func (r AlibabaScbpKeywordMatchedProductsGetAPIRequest) GetAdKeyword() string {
	return r._adKeyword
}
