package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightInsuranceOrderApply 多险种批量投保
// alitrip.flight.insurance.order.apply
//
// 多险种批量投保
func AlitripFlightInsuranceOrderApply(clt *core.SDKClient, req *flightuppc.AlitripFlightInsuranceOrderApplyAPIRequest, session string) (*flightuppc.AlitripFlightInsuranceOrderApplyAPIResponse, error) {
	var resp flightuppc.AlitripFlightInsuranceOrderApplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
