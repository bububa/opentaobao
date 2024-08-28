package flightuppc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightInsuranceOrderModify 保险订单批改申请
// alitrip.flight.insurance.order.modify
//
// 保险订单批改申请
func AlitripFlightInsuranceOrderModify(ctx context.Context, clt *core.SDKClient, req *flightuppc.AlitripFlightInsuranceOrderModifyAPIRequest, resp *flightuppc.AlitripFlightInsuranceOrderModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
