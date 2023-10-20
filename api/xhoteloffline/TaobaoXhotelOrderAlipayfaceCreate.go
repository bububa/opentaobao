package xhoteloffline

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhoteloffline"
)

// TaobaoXhotelOrderAlipayfaceCreate 信用住支付创建接口
// taobao.xhotel.order.alipayface.create
//
// 用于创建一笔信用住支付，主要应用场景是线下信用住
func TaobaoXhotelOrderAlipayfaceCreate(clt *core.SDKClient, req *xhoteloffline.TaobaoXhotelOrderAlipayfaceCreateAPIRequest, resp *xhoteloffline.TaobaoXhotelOrderAlipayfaceCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
