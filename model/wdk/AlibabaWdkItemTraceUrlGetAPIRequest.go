package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemTraceUrlGetAPIRequest
根据shopId和skuCode返回商品静态溯源url API请求
alibaba.wdk.item.trace.url.get

根据shopId和skuCode返回商品静态溯源url */
type AlibabaWdkItemTraceUrlGetAPIRequest struct {
	model.Params
	// 所属门店code
	_shopId string
	// 来源编码
	_sourceCode string
	// 商品skuCode
	_skuCode string
}

// New
