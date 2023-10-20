package consignplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/consignplatform"
)

// CainiaoConsignplatformOrderCancel 菜鸟发货工作台取消包裹以及订单
// cainiao.consignplatform.order.cancel
//
// 菜鸟发货工作台，商家或者isv通过api取消包裹、回收单号，如果是裹裹运力会取消小件员上门。最后删除订单信息。
func CainiaoConsignplatformOrderCancel(clt *core.SDKClient, req *consignplatform.CainiaoConsignplatformOrderCancelAPIRequest, resp *consignplatform.CainiaoConsignplatformOrderCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
