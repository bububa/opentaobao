package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripsupplychaintrainsearch 【商旅】火车票订单查询
// alitrip.btrip.supplychain.train.search
//
// 【商旅】火车票订单查询
func Alitripbtripsupplychaintrainsearch(clt *core.SDKClient, req *btrip.AlitripbtripsupplychaintrainsearchAPIRequest, session string) (*btrip.AlitripbtripsupplychaintrainsearchAPIResponse, error) {
	var resp btrip.AlitripbtripsupplychaintrainsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
