package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// AlitripFlightInsuranceProductSearch 搜索保险产品
// alitrip.flight.insurance.product.search
//
// 搜索保险产品
func AlitripFlightInsuranceProductSearch(clt *core.SDKClient, req *flightuppc.AlitripFlightInsuranceProductSearchAPIRequest, session string) (*flightuppc.AlitripFlightInsuranceProductSearchAPIResponse, error) {
	var resp flightuppc.AlitripFlightInsuranceProductSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
