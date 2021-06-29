package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
aliexpress.solution.seller.category.tree.query 
aliexpress.solution.seller.category.tree.query

API for seller to query the category tree. Support only displaying the categories which seller have permissions to publish products.
*/
func AliexpressSolutionSellerCategoryTreeQuery(clt *core.SDKClient, req *aesolution.AliexpressSolutionSellerCategoryTreeQueryRequest, session string) (*aesolution.AliexpressSolutionSellerCategoryTreeQueryAPIResponse, error) {
    var resp aesolution.AliexpressSolutionSellerCategoryTreeQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
