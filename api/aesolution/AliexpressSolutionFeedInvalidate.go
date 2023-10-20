package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionFeedInvalidate aliexpress.solution.feed.invalidate
// aliexpress.solution.feed.invalidate
//
// Api for invalidating specific feeds based on job Ids. Please use aliexpress.solution.feed.list.get to determine which job Ids needs to be sent for invalidation.
func AliexpressSolutionFeedInvalidate(clt *core.SDKClient, req *aesolution.AliexpressSolutionFeedInvalidateAPIRequest, resp *aesolution.AliexpressSolutionFeedInvalidateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
