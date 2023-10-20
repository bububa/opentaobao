package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcostcenterentitydelete 删除外部成本中心人员信息
// alitrip.btrip.cost.center.entity.delete
//
// 删除外部成本中心人员信息
func Alitripbtripcostcenterentitydelete(clt *core.SDKClient, req *btrip.AlitripbtripcostcenterentitydeleteAPIRequest, session string) (*btrip.AlitripbtripcostcenterentitydeleteAPIResponse, error) {
	var resp btrip.AlitripbtripcostcenterentitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
