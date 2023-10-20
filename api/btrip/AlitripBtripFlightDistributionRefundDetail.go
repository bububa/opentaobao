package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundDetail 商旅机票退票详情接口
// alitrip.btrip.flight.distribution.refund.detail
//
// 商旅机票分销退票详情
func AlitripBtripFlightDistributionRefundDetail(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundDetailAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionRefundDetailAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionRefundDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
