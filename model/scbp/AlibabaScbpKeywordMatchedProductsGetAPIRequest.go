package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpKeywordMatchedProductsGetAPIRequest
查询和词匹配的推广产品 API请求
alibaba.scbp.keyword.matched.products.get

查询和词匹配的推广产品 */
type AlibabaScbpKeywordMatchedProductsGetAPIRequest struct {
	model.Params
	// 已购买的关键词
	_adKeyword string
}

// New
