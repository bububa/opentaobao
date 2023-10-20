package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// Aliexpresssolutionfeedlistget aliexpress.solution.feed.list.get
// aliexpress.solution.feed.list.get
//
// API to query the feed list belonged to a seller
func Aliexpresssolutionfeedlistget(clt *core.SDKClient, req *aesolution.AliexpresssolutionfeedlistgetAPIRequest, session string) (*aesolution.AliexpresssolutionfeedlistgetAPIResponse, error) {
	var resp aesolution.AliexpresssolutionfeedlistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
