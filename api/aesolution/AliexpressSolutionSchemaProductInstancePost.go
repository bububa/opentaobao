package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionschemaproductinstancepost aliexpress.solution.schema.product.instance.post
// aliexpress.solution.schema.product.instance.post
//
// Upload product based on json schema instance.QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
func Aliexpresssolutionschemaproductinstancepost(clt *core.SDKClient, req *aesolution.AliexpresssolutionschemaproductinstancepostAPIRequest, session string) (*aesolution.AliexpresssolutionschemaproductinstancepostAPIResponse, error) {
	var resp aesolution.AliexpresssolutionschemaproductinstancepostAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
