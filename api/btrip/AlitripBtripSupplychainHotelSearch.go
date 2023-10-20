package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripsupplychainhotelsearch 【商旅】酒店订单查询
// alitrip.btrip.supplychain.hotel.search
//
// 【商旅】酒店订单查询
func Alitripbtripsupplychainhotelsearch(clt *core.SDKClient, req *btrip.AlitripbtripsupplychainhotelsearchAPIRequest, session string) (*btrip.AlitripbtripsupplychainhotelsearchAPIResponse, error) {
	var resp btrip.AlitripbtripsupplychainhotelsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
