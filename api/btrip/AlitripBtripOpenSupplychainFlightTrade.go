package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopensupplychainflighttrade 【商旅】机票交易流水查询接口
// alitrip.btrip.open.supplychain.flight.trade
//
// 【商旅】杭州市政府机票交易流水接口查询
func Alitripbtripopensupplychainflighttrade(clt *core.SDKClient, req *btrip.AlitripbtripopensupplychainflighttradeAPIRequest, session string) (*btrip.AlitripbtripopensupplychainflighttradeAPIResponse, error) {
	var resp btrip.AlitripbtripopensupplychainflighttradeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
