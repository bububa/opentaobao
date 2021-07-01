package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

/* AliexpressSolutionSchemaProductFullUpdate
aliexpress.solution.schema.product.full.update
aliexpress.solution.schema.product.full.update

Schema interface for product full update. QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller. */
func AliexpressSolutionSchemaProductFullUpdate(clt *core.SDKClient, req *aesolution.AliexpressSolutionSchemaProductFullUpdateAPIRequest, session string) (*aesolution.AliexpressSolutionSchemaProductFullUpdateAPIResponse, error) {
	var resp aesolution.AliexpressSolutionSchemaProductFullUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
