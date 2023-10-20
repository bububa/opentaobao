package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopensupplychaintraintrade 商旅火车票交易流水接口
// alitrip.btrip.open.supplychain.train.trade
//
// 商旅火车票交易流水接口
func Alitripbtripopensupplychaintraintrade(clt *core.SDKClient, req *btrip.AlitripbtripopensupplychaintraintradeAPIRequest, session string) (*btrip.AlitripbtripopensupplychaintraintradeAPIResponse, error) {
	var resp btrip.AlitripbtripopensupplychaintraintradeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
