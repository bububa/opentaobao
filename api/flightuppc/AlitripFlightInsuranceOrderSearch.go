package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightInsuranceOrderSearch 查询保险订单详情
// alitrip.flight.insurance.order.search
//
// 查询保险订单详情
func AlitripFlightInsuranceOrderSearch(clt *core.SDKClient, req *flightuppc.AlitripFlightInsuranceOrderSearchAPIRequest, resp *flightuppc.AlitripFlightInsuranceOrderSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
