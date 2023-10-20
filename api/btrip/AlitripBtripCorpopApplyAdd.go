package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopapplyadd 【商旅】isv添加审批单
// alitrip.btrip.corpop.apply.add
//
// 【商旅】isv添加审批单
func Alitripbtripcorpopapplyadd(clt *core.SDKClient, req *btrip.AlitripbtripcorpopapplyaddAPIRequest, session string) (*btrip.AlitripbtripcorpopapplyaddAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopapplyaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
