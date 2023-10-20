package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeShipaddressUpdate Openmall订单收货地址修改
// taobao.openmall.trade.shipaddress.update
//
// Openmall订单收货地址修改
func TaobaoOpenmallTradeShipaddressUpdate(clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeShipaddressUpdateAPIRequest, resp *openmall.TaobaoOpenmallTradeShipaddressUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
