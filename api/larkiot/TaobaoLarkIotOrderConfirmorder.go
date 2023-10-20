package larkiot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/larkiot"
)

// TaobaoLarkIotOrderConfirmorder iot渠道卖品落单
// taobao.lark.iot.order.confirmorder
//
// 云智对接无人超市，接收无人超市订单信息
func TaobaoLarkIotOrderConfirmorder(clt *core.SDKClient, req *larkiot.TaobaoLarkIotOrderConfirmorderAPIRequest, resp *larkiot.TaobaoLarkIotOrderConfirmorderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
