package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest
根据商品id获取商品属性 API请求
alibaba.alihealth.tracecodeseller.product.attr.search

根据商品id获取商品属性 */
type AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest struct {
	model.Params
	// 企业id
	_entInfoId int64
	// 货品id
	_tracUserProductInfoId int64
}

// New
