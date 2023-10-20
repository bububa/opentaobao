package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeConfirm 确认收货
// taobao.openmall.trade.confirm
//
// 确认订单收货
func TaobaoOpenmallTradeConfirm(clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeConfirmAPIRequest, resp *openmall.TaobaoOpenmallTradeConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
