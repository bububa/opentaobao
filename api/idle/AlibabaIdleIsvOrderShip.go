package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleIsvOrderShip 闲鱼无忧购服务商发货
// alibaba.idle.isv.order.ship
//
// 闲鱼无忧购业务入仓模式下服务商订单发货的接口
func AlibabaIdleIsvOrderShip(clt *core.SDKClient, req *idle.AlibabaIdleIsvOrderShipAPIRequest, session string) (*idle.AlibabaIdleIsvOrderShipAPIResponse, error) {
	var resp idle.AlibabaIdleIsvOrderShipAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
