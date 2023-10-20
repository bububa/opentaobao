package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightInsuranceProductSearch 搜索保险产品
// alitrip.flight.insurance.product.search
//
// 搜索保险产品
func AlitripFlightInsuranceProductSearch(clt *core.SDKClient, req *flightuppc.AlitripFlightInsuranceProductSearchAPIRequest, resp *flightuppc.AlitripFlightInsuranceProductSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
