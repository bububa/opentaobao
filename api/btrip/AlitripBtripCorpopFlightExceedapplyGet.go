package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopflightexceedapplyget 商旅机票第三方超标审批单搜索接口
// alitrip.btrip.corpop.flight.exceedapply.get
//
// 商旅机票第三方超标审批单搜索接口
func Alitripbtripcorpopflightexceedapplyget(clt *core.SDKClient, req *btrip.AlitripbtripcorpopflightexceedapplygetAPIRequest, session string) (*btrip.AlitripbtripcorpopflightexceedapplygetAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopflightexceedapplygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
