package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripsupplychainflightsearch 【商旅】机票订单查询
// alitrip.btrip.supplychain.flight.search
//
// 【商旅】机票订单查询
func Alitripbtripsupplychainflightsearch(clt *core.SDKClient, req *btrip.AlitripbtripsupplychainflightsearchAPIRequest, session string) (*btrip.AlitripbtripsupplychainflightsearchAPIResponse, error) {
	var resp btrip.AlitripbtripsupplychainflightsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
