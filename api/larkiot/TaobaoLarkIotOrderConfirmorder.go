package larkiot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/larkiot"
)

/* TaobaoLarkIotOrderConfirmorder
iot渠道卖品落单
taobao.lark.iot.order.confirmorder

云智对接无人超市，接收无人超市订单信息 */
func TaobaoLarkIotOrderConfirmorder(clt *core.SDKClient, req *larkiot.TaobaoLarkIotOrderConfirmorderAPIRequest, session string) (*larkiot.TaobaoLarkIotOrderConfirmorderAPIResponse, error) {
	var resp larkiot.TaobaoLarkIotOrderConfirmorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
