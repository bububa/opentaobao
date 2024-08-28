package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightOrderSearch 获取机票订单列表
// alitrip.btrip.flight.order.search
//
// 第三方OA厂商获取机票订单列表
func AlitripBtripFlightOrderSearch(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightOrderSearchAPIRequest, resp *btrip.AlitripBtripFlightOrderSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
