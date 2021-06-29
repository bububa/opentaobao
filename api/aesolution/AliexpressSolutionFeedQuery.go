package aesolution

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aesolution"
)

/* 
aliexpress.solution.feed.query 
aliexpress.solution.feed.query

API for query the execution result of feed.
*/
func AliexpressSolutionFeedQuery(clt *core.SDKClient, req *aesolution.AliexpressSolutionFeedQueryRequest, session string) (*aesolution.AliexpressSolutionFeedQueryAPIResponse, error) {
    var resp aesolution.AliexpressSolutionFeedQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}