package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripHotelDistributionOrderPay
商旅酒店分销订单支付
alitrip.btrip.hotel.distribution.order.pay

商旅酒店分销订单支付 */
func AlitripBtripHotelDistributionOrderPay(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionOrderPayAPIRequest, session string) (*btrip.AlitripBtripHotelDistributionOrderPayAPIResponse, error) {
	var resp btrip.AlitripBtripHotelDistributionOrderPayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
