package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightInsuranceOrderModify 保险订单批改申请
// alitrip.flight.insurance.order.modify
//
// 保险订单批改申请
func AlitripFlightInsuranceOrderModify(clt *core.SDKClient, req *flightuppc.AlitripFlightInsuranceOrderModifyAPIRequest, session string) (*flightuppc.AlitripFlightInsuranceOrderModifyAPIResponse, error) {
	var resp flightuppc.AlitripFlightInsuranceOrderModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
