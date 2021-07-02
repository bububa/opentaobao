package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionFeedListGet aliexpress.solution.feed.list.get
// aliexpress.solution.feed.list.get
//
// API to query the feed list belonged to a seller
func AliexpressSolutionFeedListGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionFeedListGetAPIRequest, session string) (*aesolution.AliexpressSolutionFeedListGetAPIResponse, error) {
	var resp aesolution.AliexpressSolutionFeedListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
