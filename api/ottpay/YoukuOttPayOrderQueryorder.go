package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttPayOrderQueryorder 查询订单
// youku.ott.pay.order.queryorder
//
// 通过订单号查询订单信息
func YoukuOttPayOrderQueryorder(clt *core.SDKClient, req *ottpay.YoukuOttPayOrderQueryorderAPIRequest, resp *ottpay.YoukuOttPayOrderQueryorderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
