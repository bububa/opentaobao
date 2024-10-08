package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionFeedSubmit aliexpress.solution.feed.submit
// aliexpress.solution.feed.submit
//
// API for merchants to submit feed data. Please note for each seller, the recommended number of feeds submitted for each operation_type every 24 hours should be lee than 150, otherwise significant delay might be encountered for processing the feed.
func AliexpressSolutionFeedSubmit(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionFeedSubmitAPIRequest, resp *aesolution.AliexpressSolutionFeedSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
