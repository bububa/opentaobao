package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// Alitripflightinsuranceordermodify 保险订单批改申请
// alitrip.flight.insurance.order.modify
//
// 保险订单批改申请
func Alitripflightinsuranceordermodify(clt *core.SDKClient, req *flightuppc.AlitripflightinsuranceordermodifyAPIRequest, session string) (*flightuppc.AlitripflightinsuranceordermodifyAPIResponse, error) {
	var resp flightuppc.AlitripflightinsuranceordermodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
