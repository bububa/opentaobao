package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeGet 查询订单详情
// taobao.openmall.trade.get
//
// 查询订单详情
func TaobaoOpenmallTradeGet(clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeGetAPIRequest, resp *openmall.TaobaoOpenmallTradeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
