package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionSchemaProductInstancePost aliexpress.solution.schema.product.instance.post
// aliexpress.solution.schema.product.instance.post
//
// Upload product based on json schema instance.QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
func AliexpressSolutionSchemaProductInstancePost(clt *core.SDKClient, req *aesolution.AliexpressSolutionSchemaProductInstancePostAPIRequest, resp *aesolution.AliexpressSolutionSchemaProductInstancePostAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
