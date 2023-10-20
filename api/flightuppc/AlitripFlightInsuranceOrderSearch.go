package flightuppc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flightuppc"
)

// Alitripflightinsuranceordersearch 查询保险订单详情
// alitrip.flight.insurance.order.search
//
// 查询保险订单详情
func Alitripflightinsuranceordersearch(clt *core.SDKClient, req *flightuppc.AlitripflightinsuranceordersearchAPIRequest, session string) (*flightuppc.AlitripflightinsuranceordersearchAPIResponse, error) {
	var resp flightuppc.AlitripflightinsuranceordersearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
