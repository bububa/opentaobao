package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripsupplychaintraindetailsearchvtwo 【商旅】火车票订单详情查询
// alitrip.btrip.supplychain.train.detail.search.vtwo
//
// 【商旅】火车票订单详情查询
func Alitripbtripsupplychaintraindetailsearchvtwo(clt *core.SDKClient, req *btrip.AlitripbtripsupplychaintraindetailsearchvtwoAPIRequest, session string) (*btrip.AlitripbtripsupplychaintraindetailsearchvtwoAPIResponse, error) {
	var resp btrip.AlitripbtripsupplychaintraindetailsearchvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
