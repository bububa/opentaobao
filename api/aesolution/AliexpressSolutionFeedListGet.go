package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionFeedListGet aliexpress.solution.feed.list.get
// aliexpress.solution.feed.list.get
//
// API to query the feed list belonged to a seller
func AliexpressSolutionFeedListGet(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionFeedListGetAPIRequest, resp *aesolution.AliexpressSolutionFeedListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
