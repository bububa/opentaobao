package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoFuwuPurchaseOrderPay 内购服务订单付款页获取接口
// taobao.fuwu.purchase.order.pay
//
// 通过接口获取某一订单的付款页面链接
func TaobaoFuwuPurchaseOrderPay(clt *core.SDKClient, req *servicecenter.TaobaoFuwuPurchaseOrderPayAPIRequest, session string) (*servicecenter.TaobaoFuwuPurchaseOrderPayAPIResponse, error) {
	var resp servicecenter.TaobaoFuwuPurchaseOrderPayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
