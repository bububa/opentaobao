package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainFlightSearch 【商旅】机票订单查询
// alitrip.btrip.supplychain.flight.search
//
// 【商旅】机票订单查询
func AlitripBtripSupplychainFlightSearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainFlightSearchAPIRequest, session string) (*btrip.AlitripBtripSupplychainFlightSearchAPIResponse, error) {
	var resp btrip.AlitripBtripSupplychainFlightSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
