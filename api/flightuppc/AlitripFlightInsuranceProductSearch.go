package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// Alitripflightinsuranceproductsearch 搜索保险产品
// alitrip.flight.insurance.product.search
//
// 搜索保险产品
func Alitripflightinsuranceproductsearch(clt *core.SDKClient, req *flightuppc.AlitripflightinsuranceproductsearchAPIRequest, session string) (*flightuppc.AlitripflightinsuranceproductsearchAPIResponse, error) {
	var resp flightuppc.AlitripflightinsuranceproductsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
