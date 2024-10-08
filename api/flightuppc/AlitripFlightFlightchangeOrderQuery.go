package flightuppc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightFlightchangeOrderQuery 订单维度航变查询
// alitrip.flight.flightchange.order.query
//
// 订单维度航变查询
func AlitripFlightFlightchangeOrderQuery(ctx context.Context, clt *core.SDKClient, req *flightuppc.AlitripFlightFlightchangeOrderQueryAPIRequest, resp *flightuppc.AlitripFlightFlightchangeOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
