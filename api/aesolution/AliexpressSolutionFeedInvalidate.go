package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionFeedInvalidate aliexpress.solution.feed.invalidate
// aliexpress.solution.feed.invalidate
//
// Api for invalidating specific feeds based on job Ids. Please use aliexpress.solution.feed.list.get to determine which job Ids needs to be sent for invalidation.
func AliexpressSolutionFeedInvalidate(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionFeedInvalidateAPIRequest, resp *aesolution.AliexpressSolutionFeedInvalidateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
