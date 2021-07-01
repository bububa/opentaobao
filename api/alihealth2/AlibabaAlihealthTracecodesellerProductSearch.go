package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaAlihealthTracecodesellerProductSearch
查询商品api
alibaba.alihealth.tracecodeseller.product.search

查询商品api */
func AlibabaAlihealthTracecodesellerProductSearch(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerProductSearchAPIRequest, session string) (*alihealth2.AlibabaAlihealthTracecodesellerProductSearchAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthTracecodesellerProductSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
