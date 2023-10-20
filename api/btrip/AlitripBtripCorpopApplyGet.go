package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopapplyget 【商旅】查询审批单
// alitrip.btrip.corpop.apply.get
//
// 【商旅】查询审批单
func Alitripbtripcorpopapplyget(clt *core.SDKClient, req *btrip.AlitripbtripcorpopapplygetAPIRequest, session string) (*btrip.AlitripbtripcorpopapplygetAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopapplygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
