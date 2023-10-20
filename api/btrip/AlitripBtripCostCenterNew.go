package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcostcenternew 新建外部成本中心
// alitrip.btrip.cost.center.new
//
// 新建外部成本中心
func Alitripbtripcostcenternew(clt *core.SDKClient, req *btrip.AlitripbtripcostcenternewAPIRequest, session string) (*btrip.AlitripbtripcostcenternewAPIResponse, error) {
	var resp btrip.AlitripbtripcostcenternewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
