package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcostcenterdelete 删除外部成本中心
// alitrip.btrip.cost.center.delete
//
// 删除外部成本中心
func Alitripbtripcostcenterdelete(clt *core.SDKClient, req *btrip.AlitripbtripcostcenterdeleteAPIRequest, session string) (*btrip.AlitripbtripcostcenterdeleteAPIResponse, error) {
	var resp btrip.AlitripbtripcostcenterdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
