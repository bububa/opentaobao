package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// Alitripflightinsuranceorderrefund 保险订单退保
// alitrip.flight.insurance.order.refund
//
// 保险订单退保
func Alitripflightinsuranceorderrefund(clt *core.SDKClient, req *flightuppc.AlitripflightinsuranceorderrefundAPIRequest, session string) (*flightuppc.AlitripflightinsuranceorderrefundAPIResponse, error) {
	var resp flightuppc.AlitripflightinsuranceorderrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
