package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionschemaproductfullupdate aliexpress.solution.schema.product.full.update
// aliexpress.solution.schema.product.full.update
//
// Schema interface for product full update. QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
func Aliexpresssolutionschemaproductfullupdate(clt *core.SDKClient, req *aesolution.AliexpresssolutionschemaproductfullupdateAPIRequest, session string) (*aesolution.AliexpresssolutionschemaproductfullupdateAPIResponse, error) {
	var resp aesolution.AliexpresssolutionschemaproductfullupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
