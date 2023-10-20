package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundNegotiatereturnRender 协商退货退款渲染
// taobao.refund.negotiatereturn.render
//
// 协商退货退款渲染
func TaobaoRefundNegotiatereturnRender(clt *core.SDKClient, req *tbrefund.TaobaoRefundNegotiatereturnRenderAPIRequest, resp *tbrefund.TaobaoRefundNegotiatereturnRenderAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
