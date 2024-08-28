package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainFlightSearch 【商旅】机票订单查询
// alitrip.btrip.supplychain.flight.search
//
// 【商旅】机票订单查询
func AlitripBtripSupplychainFlightSearch(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripSupplychainFlightSearchAPIRequest, resp *btrip.AlitripBtripSupplychainFlightSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
