package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtriphoteldistributionorderpay 商旅酒店分销订单支付
// alitrip.btrip.hotel.distribution.order.pay
//
// 商旅酒店分销订单支付
func Alitripbtriphoteldistributionorderpay(clt *core.SDKClient, req *btrip.AlitripbtriphoteldistributionorderpayAPIRequest, session string) (*btrip.AlitripbtriphoteldistributionorderpayAPIResponse, error) {
	var resp btrip.AlitripbtriphoteldistributionorderpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
