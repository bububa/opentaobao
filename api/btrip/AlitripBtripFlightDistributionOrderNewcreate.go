package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionOrderNewcreate 商旅机票分销-创建订单V2
// alitrip.btrip.flight.distribution.order.newcreate
//
// 商旅机票分销-创建订单V2
func AlitripBtripFlightDistributionOrderNewcreate(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionOrderNewcreateAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionOrderNewcreateAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionOrderNewcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
