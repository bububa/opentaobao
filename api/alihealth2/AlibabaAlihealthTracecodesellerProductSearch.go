package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthtracecodesellerproductsearch 查询商品api
// alibaba.alihealth.tracecodeseller.product.search
//
// 查询商品api
func Alibabaalihealthtracecodesellerproductsearch(clt *core.SDKClient, req *alihealth2.AlibabaalihealthtracecodesellerproductsearchAPIRequest, session string) (*alihealth2.AlibabaalihealthtracecodesellerproductsearchAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthtracecodesellerproductsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
