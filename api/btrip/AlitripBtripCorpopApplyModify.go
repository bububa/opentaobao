package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopapplymodify 【商旅】修改出差审批单（行程）
// alitrip.btrip.corpop.apply.modify
//
// 【商旅】修改出差审批单（行程）
func Alitripbtripcorpopapplymodify(clt *core.SDKClient, req *btrip.AlitripbtripcorpopapplymodifyAPIRequest, session string) (*btrip.AlitripbtripcorpopapplymodifyAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopapplymodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
