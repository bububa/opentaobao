package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightFlightchangeOrderQuery 订单维度航变查询
// alitrip.flight.flightchange.order.query
//
// 订单维度航变查询
func AlitripFlightFlightchangeOrderQuery(clt *core.SDKClient, req *flightuppc.AlitripFlightFlightchangeOrderQueryAPIRequest, session string) (*flightuppc.AlitripFlightFlightchangeOrderQueryAPIResponse, error) {
	var resp flightuppc.AlitripFlightFlightchangeOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
