package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// Youkuottpayorderqueryauthstate 查询连包签约状态
// youku.ott.pay.order.queryauthstate
//
// 查询CP用户连包商品签约状态
func Youkuottpayorderqueryauthstate(clt *core.SDKClient, req *ottpay.YoukuottpayorderqueryauthstateAPIRequest, session string) (*ottpay.YoukuottpayorderqueryauthstateAPIResponse, error) {
	var resp ottpay.YoukuottpayorderqueryauthstateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
