package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcostcenterentityadd 增加外部成本中心人员信息
// alitrip.btrip.cost.center.entity.add
//
// 增加外部成本中心人员信息
func Alitripbtripcostcenterentityadd(clt *core.SDKClient, req *btrip.AlitripbtripcostcenterentityaddAPIRequest, session string) (*btrip.AlitripbtripcostcenterentityaddAPIResponse, error) {
	var resp btrip.AlitripbtripcostcenterentityaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
