package flightuppc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightInsuranceOrderApply 多险种批量投保
// alitrip.flight.insurance.order.apply
//
// 多险种批量投保
func AlitripFlightInsuranceOrderApply(ctx context.Context, clt *core.SDKClient, req *flightuppc.AlitripFlightInsuranceOrderApplyAPIRequest, resp *flightuppc.AlitripFlightInsuranceOrderApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
