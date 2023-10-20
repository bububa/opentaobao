package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthtracecodesellerproductattrsearch 根据商品id获取商品属性
// alibaba.alihealth.tracecodeseller.product.attr.search
//
// 根据商品id获取商品属性
func Alibabaalihealthtracecodesellerproductattrsearch(clt *core.SDKClient, req *drugtrace.AlibabaalihealthtracecodesellerproductattrsearchAPIRequest, session string) (*drugtrace.AlibabaalihealthtracecodesellerproductattrsearchAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthtracecodesellerproductattrsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
