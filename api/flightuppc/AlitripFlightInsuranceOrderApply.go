package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// Alitripflightinsuranceorderapply 多险种批量投保
// alitrip.flight.insurance.order.apply
//
// 多险种批量投保
func Alitripflightinsuranceorderapply(clt *core.SDKClient, req *flightuppc.AlitripflightinsuranceorderapplyAPIRequest, session string) (*flightuppc.AlitripflightinsuranceorderapplyAPIResponse, error) {
	var resp flightuppc.AlitripflightinsuranceorderapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
