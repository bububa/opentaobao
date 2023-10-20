package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopapplyapprove 【商旅】更新审批单状态
// alitrip.btrip.corpop.apply.approve
//
// 【商旅】更新审批单状态
func Alitripbtripcorpopapplyapprove(clt *core.SDKClient, req *btrip.AlitripbtripcorpopapplyapproveAPIRequest, session string) (*btrip.AlitripbtripcorpopapplyapproveAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopapplyapproveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
