package larkiot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/larkiot"
)

// Taobaolarkiotorderconfirmorder iot渠道卖品落单
// taobao.lark.iot.order.confirmorder
//
// 云智对接无人超市，接收无人超市订单信息
func Taobaolarkiotorderconfirmorder(clt *core.SDKClient, req *larkiot.TaobaolarkiotorderconfirmorderAPIRequest, session string) (*larkiot.TaobaolarkiotorderconfirmorderAPIResponse, error) {
	var resp larkiot.TaobaolarkiotorderconfirmorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
