package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

/* AliexpressSolutionProductListGet
Get product list
aliexpress.solution.product.list.get

Get product list */
func AliexpressSolutionProductListGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionProductListGetAPIRequest, session string) (*aesolution.AliexpressSolutionProductListGetAPIResponse, error) {
	var resp aesolution.AliexpressSolutionProductListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
