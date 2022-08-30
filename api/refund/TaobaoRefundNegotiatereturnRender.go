package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// TaobaoRefundNegotiatereturnRender 协商退货退款渲染
// taobao.refund.negotiatereturn.render
//
// 协商退货退款渲染
func TaobaoRefundNegotiatereturnRender(clt *core.SDKClient, req *refund.TaobaoRefundNegotiatereturnRenderAPIRequest, session string) (*refund.TaobaoRefundNegotiatereturnRenderAPIResponse, error) {
	var resp refund.TaobaoRefundNegotiatereturnRenderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
