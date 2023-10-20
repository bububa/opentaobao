package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopcommonapplyget 商旅审批单通用查询接口
// alitrip.btrip.corpop.commonapply.get
//
// 商旅审批单通用查询接口
func Alitripbtripcorpopcommonapplyget(clt *core.SDKClient, req *btrip.AlitripbtripcorpopcommonapplygetAPIRequest, session string) (*btrip.AlitripbtripcorpopcommonapplygetAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopcommonapplygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
