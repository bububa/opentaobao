package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripcitycarapplyquery 三方市内用车申请单查询
// alitrip.btrip.city.car.apply.query
//
// 三方市内用车申请单查询
func Alitripbtripcitycarapplyquery(clt *core.SDKClient, req *btrip.AlitripbtripcitycarapplyqueryAPIRequest, session string) (*btrip.AlitripbtripcitycarapplyqueryAPIResponse, error) {
	var resp btrip.AlitripbtripcitycarapplyqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
