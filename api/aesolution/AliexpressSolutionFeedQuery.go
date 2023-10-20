package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionFeedQuery aliexpress.solution.feed.query
// aliexpress.solution.feed.query
//
// API for query the execution result of feed.
func AliexpressSolutionFeedQuery(clt *core.SDKClient, req *aesolution.AliexpressSolutionFeedQueryAPIRequest, resp *aesolution.AliexpressSolutionFeedQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
