package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// AlibabaOrderPayResultQuery alibaba查询订单支付结果
// alibaba.order.pay.result.query
//
// alibaba查询订单支付结果
func AlibabaOrderPayResultQuery(clt *core.SDKClient, req *icbudropshipping.AlibabaOrderPayResultQueryAPIRequest, resp *icbudropshipping.AlibabaOrderPayResultQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
