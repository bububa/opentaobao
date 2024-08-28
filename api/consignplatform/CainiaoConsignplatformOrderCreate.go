package consignplatform

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/consignplatform"
)

// CainiaoConsignplatformOrderCreate 菜鸟发货工作台创建订单
// cainiao.consignplatform.order.create
//
// 菜鸟发货工作台，商家或者isv通过api进行订单写入操作
func CainiaoConsignplatformOrderCreate(ctx context.Context, clt *core.SDKClient, req *consignplatform.CainiaoConsignplatformOrderCreateAPIRequest, resp *consignplatform.CainiaoConsignplatformOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
