package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcostcenterentityset 设置外部成本中心人员信息
// alitrip.btrip.cost.center.entity.set
//
// 设置外部成本中心人员信息
func Alitripbtripcostcenterentityset(clt *core.SDKClient, req *btrip.AlitripbtripcostcenterentitysetAPIRequest, session string) (*btrip.AlitripbtripcostcenterentitysetAPIResponse, error) {
	var resp btrip.AlitripbtripcostcenterentitysetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
