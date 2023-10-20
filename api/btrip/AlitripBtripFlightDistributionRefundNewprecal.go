package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundNewprecal 商旅机票分销-退票费预计算
// alitrip.btrip.flight.distribution.refund.newprecal
//
// 商旅机票分销-退票费预计算
func AlitripBtripFlightDistributionRefundNewprecal(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundNewprecalAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionRefundNewprecalAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionRefundNewprecalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
