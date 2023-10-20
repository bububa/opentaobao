package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionfeedquery aliexpress.solution.feed.query
// aliexpress.solution.feed.query
//
// API for query the execution result of feed.
func Aliexpresssolutionfeedquery(clt *core.SDKClient, req *aesolution.AliexpresssolutionfeedqueryAPIRequest, session string) (*aesolution.AliexpresssolutionfeedqueryAPIResponse, error) {
	var resp aesolution.AliexpresssolutionfeedqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
