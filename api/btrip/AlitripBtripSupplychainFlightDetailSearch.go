package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainFlightDetailSearch 【商旅】机票订单详情查询
// alitrip.btrip.supplychain.flight.detail.search
//
// 【商旅】机票订单详情查询
func AlitripBtripSupplychainFlightDetailSearch(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripSupplychainFlightDetailSearchAPIRequest, resp *btrip.AlitripBtripSupplychainFlightDetailSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
