package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopensupplychainhoteltrade 【商旅】酒店交易查询流水接口
// alitrip.btrip.open.supplychain.hotel.trade
//
// 【商旅】酒店交易查询流水接口——杭州市政府
func Alitripbtripopensupplychainhoteltrade(clt *core.SDKClient, req *btrip.AlitripbtripopensupplychainhoteltradeAPIRequest, session string) (*btrip.AlitripbtripopensupplychainhoteltradeAPIResponse, error) {
	var resp btrip.AlitripbtripopensupplychainhoteltradeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
