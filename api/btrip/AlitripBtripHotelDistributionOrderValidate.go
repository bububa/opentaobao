package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionOrderValidate 商旅酒店API分销下单前校验
// alitrip.btrip.hotel.distribution.order.validate
//
// 商旅酒店API分销下单前校验
func AlitripBtripHotelDistributionOrderValidate(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionOrderValidateAPIRequest, resp *btrip.AlitripBtripHotelDistributionOrderValidateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
