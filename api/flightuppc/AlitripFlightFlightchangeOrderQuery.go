package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// Alitripflightflightchangeorderquery 订单维度航变查询
// alitrip.flight.flightchange.order.query
//
// 订单维度航变查询
func Alitripflightflightchangeorderquery(clt *core.SDKClient, req *flightuppc.AlitripflightflightchangeorderqueryAPIRequest, session string) (*flightuppc.AlitripflightflightchangeorderqueryAPIResponse, error) {
	var resp flightuppc.AlitripflightflightchangeorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
