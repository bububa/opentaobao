package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcorpopapplysearch 【商旅】搜索审批单列表
// alitrip.btrip.corpop.apply.search
//
// 【商旅】搜索审批单列表
func Alitripbtripcorpopapplysearch(clt *core.SDKClient, req *btrip.AlitripbtripcorpopapplysearchAPIRequest, session string) (*btrip.AlitripbtripcorpopapplysearchAPIResponse, error) {
	var resp btrip.AlitripbtripcorpopapplysearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
