package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpoptrainexceedapplyget 商旅火车票第三方超标审批单搜索接口
// alitrip.btrip.corpop.train.exceedapply.get
//
// 商旅火车票第三方超标审批单搜索接口
func Alitripbtripcorpoptrainexceedapplyget(clt *core.SDKClient, req *btrip.AlitripbtripcorpoptrainexceedapplygetAPIRequest, session string) (*btrip.AlitripbtripcorpoptrainexceedapplygetAPIResponse, error) {
	var resp btrip.AlitripbtripcorpoptrainexceedapplygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
