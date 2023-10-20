package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttPayOrderQueryauthstate 查询连包签约状态
// youku.ott.pay.order.queryauthstate
//
// 查询CP用户连包商品签约状态
func YoukuOttPayOrderQueryauthstate(clt *core.SDKClient, req *ottpay.YoukuOttPayOrderQueryauthstateAPIRequest, session string) (*ottpay.YoukuOttPayOrderQueryauthstateAPIResponse, error) {
	var resp ottpay.YoukuOttPayOrderQueryauthstateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
