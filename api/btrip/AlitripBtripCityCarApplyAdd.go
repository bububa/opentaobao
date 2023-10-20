package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcitycarapplyadd 三方市内用车申请单同步
// alitrip.btrip.city.car.apply.add
//
// 三方市内用车申请单同步
func Alitripbtripcitycarapplyadd(clt *core.SDKClient, req *btrip.AlitripbtripcitycarapplyaddAPIRequest, session string) (*btrip.AlitripbtripcitycarapplyaddAPIResponse, error) {
	var resp btrip.AlitripbtripcitycarapplyaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
