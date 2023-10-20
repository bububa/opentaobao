package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtriphoteldistributionordercancel 商旅酒店API分销取消订单
// alitrip.btrip.hotel.distribution.order.cancel
//
// 商旅酒店API分销取消订单
func Alitripbtriphoteldistributionordercancel(clt *core.SDKClient, req *btrip.AlitripbtriphoteldistributionordercancelAPIRequest, session string) (*btrip.AlitripbtriphoteldistributionordercancelAPIResponse, error) {
	var resp btrip.AlitripbtriphoteldistributionordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
