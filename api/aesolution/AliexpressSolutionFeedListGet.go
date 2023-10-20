package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionFeedListGet aliexpress.solution.feed.list.get
// aliexpress.solution.feed.list.get
//
// API to query the feed list belonged to a seller
func AliexpressSolutionFeedListGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionFeedListGetAPIRequest, resp *aesolution.AliexpressSolutionFeedListGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
