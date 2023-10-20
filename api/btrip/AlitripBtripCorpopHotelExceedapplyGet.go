package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpophotelexceedapplyget 商旅酒店第三方超标审批单搜索接口
// alitrip.btrip.corpop.hotel.exceedapply.get
//
// 商旅酒店第三方超标审批单搜索接口
func Alitripbtripcorpophotelexceedapplyget(clt *core.SDKClient, req *btrip.AlitripbtripcorpophotelexceedapplygetAPIRequest, session string) (*btrip.AlitripbtripcorpophotelexceedapplygetAPIResponse, error) {
	var resp btrip.AlitripbtripcorpophotelexceedapplygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
