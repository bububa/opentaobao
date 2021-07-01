package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

/* AliexpressSolutionFeedInvalidate
aliexpress.solution.feed.invalidate
aliexpress.solution.feed.invalidate

Api for invalidating specific feeds based on job Ids. Please use aliexpress.solution.feed.list.get to determine which job Ids needs to be sent for invalidation. */
func AliexpressSolutionFeedInvalidate(clt *core.SDKClient, req *aesolution.AliexpressSolutionFeedInvalidateAPIRequest, session string) (*aesolution.AliexpressSolutionFeedInvalidateAPIResponse, error) {
	var resp aesolution.AliexpressSolutionFeedInvalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
