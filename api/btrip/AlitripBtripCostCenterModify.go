package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcostcentermodify 修改外部成本中心
// alitrip.btrip.cost.center.modify
//
// 修改外部成本中心，设置成员，设置支付宝账号，设置名称，编号等
func Alitripbtripcostcentermodify(clt *core.SDKClient, req *btrip.AlitripbtripcostcentermodifyAPIRequest, session string) (*btrip.AlitripbtripcostcentermodifyAPIResponse, error) {
	var resp btrip.AlitripbtripcostcentermodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
