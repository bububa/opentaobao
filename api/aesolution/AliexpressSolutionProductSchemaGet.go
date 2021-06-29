package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
get product schema 
aliexpress.solution.product.schema.get

provide a new schema way to post product. With a pair of API, one for getting schema, one for posting instance
*/
func AliexpressSolutionProductSchemaGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionProductSchemaGetRequest, session string) (*aesolution.AliexpressSolutionProductSchemaGetAPIResponse, error) {
    var resp aesolution.AliexpressSolutionProductSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
