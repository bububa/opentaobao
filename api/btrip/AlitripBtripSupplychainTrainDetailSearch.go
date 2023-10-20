package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripsupplychaintraindetailsearch 【商旅】火车票订单详情查询
// alitrip.btrip.supplychain.train.detail.search
//
// 【商旅】火车票订单详情查询
func Alitripbtripsupplychaintraindetailsearch(clt *core.SDKClient, req *btrip.AlitripbtripsupplychaintraindetailsearchAPIRequest, session string) (*btrip.AlitripbtripsupplychaintraindetailsearchAPIResponse, error) {
	var resp btrip.AlitripbtripsupplychaintraindetailsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
