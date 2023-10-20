package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeRender 渲染订单价格
// taobao.openmall.trade.render
//
// 请求渲染订单价格
func TaobaoOpenmallTradeRender(clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeRenderAPIRequest, resp *openmall.TaobaoOpenmallTradeRenderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
