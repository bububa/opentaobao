package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcostcenterget 获取用户费用归属
// alitrip.btrip.cost.center.get
//
// 获取差旅申请用户的费用归属列表
func Alitripbtripcostcenterget(clt *core.SDKClient, req *btrip.AlitripbtripcostcentergetAPIRequest, session string) (*btrip.AlitripbtripcostcentergetAPIResponse, error) {
	var resp btrip.AlitripbtripcostcentergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
