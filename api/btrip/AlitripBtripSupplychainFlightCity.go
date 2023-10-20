package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripsupplychainflightcity 机场数据查询
// alitrip.btrip.supplychain.flight.city
//
// 机场数据查询
func Alitripbtripsupplychainflightcity(clt *core.SDKClient, req *btrip.AlitripbtripsupplychainflightcityAPIRequest, session string) (*btrip.AlitripbtripsupplychainflightcityAPIResponse, error) {
	var resp btrip.AlitripbtripsupplychainflightcityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
