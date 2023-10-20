package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightInsuranceOrderApply 多险种批量投保
// alitrip.flight.insurance.order.apply
//
// 多险种批量投保
func AlitripFlightInsuranceOrderApply(clt *core.SDKClient, req *flightuppc.AlitripFlightInsuranceOrderApplyAPIRequest, resp *flightuppc.AlitripFlightInsuranceOrderApplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
