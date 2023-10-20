package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripsupplychainflightdetailsearch 【商旅】机票订单详情查询
// alitrip.btrip.supplychain.flight.detail.search
//
// 【商旅】机票订单详情查询
func Alitripbtripsupplychainflightdetailsearch(clt *core.SDKClient, req *btrip.AlitripbtripsupplychainflightdetailsearchAPIRequest, session string) (*btrip.AlitripbtripsupplychainflightdetailsearchAPIResponse, error) {
	var resp btrip.AlitripbtripsupplychainflightdetailsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
