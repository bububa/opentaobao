package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtriphoteldistributionordercreate 商旅酒店分销-创建订单
// alitrip.btrip.hotel.distribution.order.create
//
// 商旅酒店分销-创建订单
func Alitripbtriphoteldistributionordercreate(clt *core.SDKClient, req *btrip.AlitripbtriphoteldistributionordercreateAPIRequest, session string) (*btrip.AlitripbtriphoteldistributionordercreateAPIResponse, error) {
	var resp btrip.AlitripbtriphoteldistributionordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
