package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthTracecodesellerProductAttrSearch 根据商品id获取商品属性
// alibaba.alihealth.tracecodeseller.product.attr.search
//
// 根据商品id获取商品属性
func AlibabaAlihealthTracecodesellerProductAttrSearch(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTracecodesellerProductAttrSearchAPIRequest, resp *drugtrace.AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
