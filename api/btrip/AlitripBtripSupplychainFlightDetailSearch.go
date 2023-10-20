package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainFlightDetailSearch 【商旅】机票订单详情查询
// alitrip.btrip.supplychain.flight.detail.search
//
// 【商旅】机票订单详情查询
func AlitripBtripSupplychainFlightDetailSearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainFlightDetailSearchAPIRequest, session string) (*btrip.AlitripBtripSupplychainFlightDetailSearchAPIResponse, error) {
	var resp btrip.AlitripBtripSupplychainFlightDetailSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
