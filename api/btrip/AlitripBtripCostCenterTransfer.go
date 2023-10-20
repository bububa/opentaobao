package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcostcentertransfer 商旅成本中心转换为外部成本中心
// alitrip.btrip.cost.center.transfer
//
// 商旅成本中心转换为外部成本中心
func Alitripbtripcostcentertransfer(clt *core.SDKClient, req *btrip.AlitripbtripcostcentertransferAPIRequest, session string) (*btrip.AlitripbtripcostcentertransferAPIResponse, error) {
	var resp btrip.AlitripbtripcostcentertransferAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
