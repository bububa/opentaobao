package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionfeedinvalidate aliexpress.solution.feed.invalidate
// aliexpress.solution.feed.invalidate
//
// Api for invalidating specific feeds based on job Ids. Please use aliexpress.solution.feed.list.get to determine which job Ids needs to be sent for invalidation.
func Aliexpresssolutionfeedinvalidate(clt *core.SDKClient, req *aesolution.AliexpresssolutionfeedinvalidateAPIRequest, session string) (*aesolution.AliexpresssolutionfeedinvalidateAPIResponse, error) {
	var resp aesolution.AliexpresssolutionfeedinvalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
