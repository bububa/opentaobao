package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoFuwuPurchaseOrderConfirm 服务市场内购服务下单接口
// taobao.fuwu.purchase.order.confirm
//
// 通过传入服务市场商品的itemcode等信息，返回给服务商内购服务的下单链接
func TaobaoFuwuPurchaseOrderConfirm(clt *core.SDKClient, req *servicecenter.TaobaoFuwuPurchaseOrderConfirmAPIRequest, resp *servicecenter.TaobaoFuwuPurchaseOrderConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
