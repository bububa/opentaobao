package consignplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/consignplatform"
)

// CainiaoConsignplatformOrderCreate 菜鸟发货工作台创建订单
// cainiao.consignplatform.order.create
//
// 菜鸟发货工作台，商家或者isv通过api进行订单写入操作
func CainiaoConsignplatformOrderCreate(clt *core.SDKClient, req *consignplatform.CainiaoConsignplatformOrderCreateAPIRequest, resp *consignplatform.CainiaoConsignplatformOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
