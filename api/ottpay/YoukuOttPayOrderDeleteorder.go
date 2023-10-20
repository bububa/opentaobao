package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttPayOrderDeleteorder 退订应用中心支付订单
// youku.ott.pay.order.deleteorder
//
// 应用中心sdk连续包月退订接口
func YoukuOttPayOrderDeleteorder(clt *core.SDKClient, req *ottpay.YoukuOttPayOrderDeleteorderAPIRequest, resp *ottpay.YoukuOttPayOrderDeleteorderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
