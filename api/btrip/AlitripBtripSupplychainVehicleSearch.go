package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripsupplychainvehiclesearch 【商旅】用车订单搜索
// alitrip.btrip.supplychain.vehicle.search
//
// 【商旅】用车订单搜索
func Alitripbtripsupplychainvehiclesearch(clt *core.SDKClient, req *btrip.AlitripbtripsupplychainvehiclesearchAPIRequest, session string) (*btrip.AlitripbtripsupplychainvehiclesearchAPIResponse, error) {
	var resp btrip.AlitripbtripsupplychainvehiclesearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
