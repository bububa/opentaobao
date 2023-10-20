package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopensupplychainvehicletrade 商旅用车交易流水接口
// alitrip.btrip.open.supplychain.vehicle.trade
//
// 商旅用车交易流水接口
func Alitripbtripopensupplychainvehicletrade(clt *core.SDKClient, req *btrip.AlitripbtripopensupplychainvehicletradeAPIRequest, session string) (*btrip.AlitripbtripopensupplychainvehicletradeAPIResponse, error) {
	var resp btrip.AlitripbtripopensupplychainvehicletradeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
