package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtriphoteldistributionordervalidate 商旅酒店API分销下单前校验
// alitrip.btrip.hotel.distribution.order.validate
//
// 商旅酒店API分销下单前校验
func Alitripbtriphoteldistributionordervalidate(clt *core.SDKClient, req *btrip.AlitripbtriphoteldistributionordervalidateAPIRequest, session string) (*btrip.AlitripbtriphoteldistributionordervalidateAPIResponse, error) {
	var resp btrip.AlitripbtriphoteldistributionordervalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
