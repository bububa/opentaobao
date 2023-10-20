package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcitycarapplyapprove 三方市内用车申请单审批
// alitrip.btrip.city.car.apply.approve
//
// 三方市内用车申请单审批
func Alitripbtripcitycarapplyapprove(clt *core.SDKClient, req *btrip.AlitripbtripcitycarapplyapproveAPIRequest, session string) (*btrip.AlitripbtripcitycarapplyapproveAPIResponse, error) {
	var resp btrip.AlitripbtripcitycarapplyapproveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
