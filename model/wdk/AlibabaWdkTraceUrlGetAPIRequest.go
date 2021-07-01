package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkTraceUrlGetAPIRequest
溯源url透出 API请求
alibaba.wdk.trace.url.get

根据shopId和skuCode返回商品溯源url */
type AlibabaWdkTraceUrlGetAPIRequest struct {
	model.Params
	// 所属门店code
	_shopId string
	// 来源编码
	_sourceCode string
	// barCode 或者skuCode
	_scanCode string
}

// New
