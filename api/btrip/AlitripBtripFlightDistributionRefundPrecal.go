package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundPrecal 商旅机票分销-退票费预计算
// alitrip.btrip.flight.distribution.refund.precal
//
// 商旅机票分销-退票费预计算
func AlitripBtripFlightDistributionRefundPrecal(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundPrecalAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionRefundPrecalAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionRefundPrecalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
